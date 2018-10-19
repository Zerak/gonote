package main

import (
	"math"
	"time"
)

type Pixel struct {
	color        int
	priceCounter int
	owner        int
}

func (p *Pixel) IsBlank() bool { return p.owner == 0 }

func (p *Pixel) CurrentPrice() int {
	if p.IsBlank() {
		return 0
	}

	return DEFAULT_PRICE * int(math.Pow(PRICE_MULTIPLIER, float64(p.priceCounter)))
}

func (p *Pixel) NextPrice() int {
	if p.IsBlank() {
		return DEFAULT_PRICE
	}

	// Put a ceiling price on Pixel. Prevents weird overflow.
	// FIXME: Unconfirmed bug. From the frontend integration test, it's possible
	// to bid up a Pixel, and the total payout of the whole game is slightly
	// more.
	//
	// At 100+ EOS per Pixel, the excess payout is ~0.006 EOS. But the team gets
	// paid ~14 EOS, so we can cover that excess. Attacker cannot make money
	// this way tho.
	//
	// Just to maintain sanity, we artificially limit the ceiling of a Pixel
	// price to below 90 EOS.
	if p.priceCounter > 25 {
		// todo
		return 9999999999999 * 10000
	}

	return int(float64(p.CurrentPrice()) * PRICE_MULTIPLIER)
}

func (p *Pixel) NextPriceCounter() int {
	if p.IsBlank() {
		return 0
	}

	return p.priceCounter + 1
}

type Account struct {
	owner int

	// Including refunds, earnings, and referral fees (scaled)
	balanceScaled int

	pixelsDrawn int
	maskScaled  int
}

func (a *Account) PrimaryKey() int { return a.owner }

type Guard struct {
	id int
	// Admin needs to refill this quota to enable more withdrwal
	quota int
}

func (g *Guard) PrimaryKey() int { return g.id }

type Canvas struct {
	id            int
	lastPaintedAt time.Time
	duration      time.Duration
	lastPainter   int
	pixelsDrawn   int

	maskScaled int
	potScaled  int
	teamScaled int
}

func (c *Canvas) PrimaryKey() int { return c.id }

func (c *Canvas) IsEnded() bool {
	return time.Now().Sub(c.lastPaintedAt.Add(c.duration)) > 0
}

func (c *Canvas) PatronBonusScaled(player *Account) int {
	return c.maskScaled*player.pixelsDrawn - player.maskScaled
}

// PixelRow pr
type PixelRow struct {
	row    int
	pixels []Pixel
}

func (p *PixelRow) PrimaryKey() int { return p.row }

func (p *PixelRow) InitializeEmptyPixels() {
	p.pixels = make([]Pixel, PIXELS_PER_ROW)
}

// PixelLocation The location of the pixel in table.
type PixelLocation struct {
	// index into the table
	row int
	// index into the pixels array of the row
	col int
}

func NewPixelLocation(coordinate int) *PixelLocation {
	return &PixelLocation{
		row: coordinate / PIXELS_PER_ROW,
		col: coordinate % PIXELS_PER_ROW,
	}
}

type PixelOrder struct {
	coordinate   int
	color        int
	priceCounter int

	x int
	y int
}

func (p *PixelOrder) Parse(memo *string) {
	//memoInt := stoull(memo, 0, 36)
	memoInt := 0 // todo
	p.priceCounter = memoInt >> 52

	p.coordinate = memoInt >> 32 & 0xFFFFF
	p.color = memoInt & 0xFFFFFFFF

	p.x = p.coordinate & 0x3ff
	p.y = p.coordinate >> 10

	//eosio_assert(y < MAX_COORDINATE_Y_PLUS_ONE, "invalid y");
	//eosio_assert(x < MAX_COORDINATE_X_PLUS_ONE, "invalid x");
}

func (p *PixelOrder) Location() *PixelLocation {
	return NewPixelLocation(p.coordinate)
}

type BuyPixelResult struct {
	// Whether the buy was skipped for some reason.
	isSkipped bool
	// Whether the buy was a blank pixel
	isFirstBuyer bool

	// Fee (scaled) paid to the contract. Going to community pot, patron bonus,
	// and team.
	feeScaled int

	// Value (scaled) that goes to the previous pixel owner. 0, if pixel was
	// blank.
	ownerEarningScaled int
}

// TransferContext Used to bookeep various money calculations
type TransferContext struct {
	purchaser int
	referrer  int

	// Fund available for purchasing
	amountLeft int // 1 EOS = 10,000

	// Total fees collected (scaled)
	totalFeesScaled int

	// How much is paid to the previous owners
	// uint128_t totalPaidToPreviousOwnersScaled;

	// How many pixels are painted
	paintedPixelCount int

	patronBonusScaled     int
	potScaled             int
	teamScaled            int
	referralEarningScaled int

	canvasMaskScaled    int
	bonusPerPixelScaled int
}

func (t *TransferContext) amountLeftScaled() int { return t.amountLeft * PRECISION_BASE }

func (t *TransferContext) purchase(p *Pixel, po *PixelOrder) BuyPixelResult {
	result := BuyPixelResult{}

	isBlank := p.IsBlank()
	result.isFirstBuyer = isBlank

	if !isBlank && po.color == p.color {
		// Pixel already the same color. Skip.
		result.isSkipped = true
		return result
	}

	if !isBlank && po.priceCounter < p.NextPriceCounter() {
		// Payment too low for this pixel, maybe price had increased (somebody
		// else drew first). Skip.
		result.isSkipped = true
		return result
	}

	nextPrice := p.NextPrice()
	//eosio_assert(amountLeft >= nextPrice, "insufficient fund to buy pixel");

	if isBlank {
		// buying blank. The fee is the entire buy price.
		result.feeScaled = nextPrice * PRECISION_BASE
	} else {
		// buying from another player. The fee is a percentage of the gain.
		currentPrice := p.CurrentPrice()
		priceIncreaseScaled := (nextPrice - currentPrice) * PRECISION_BASE

		result.feeScaled = priceIncreaseScaled * FEE_PERCENTAGE / 100
		result.ownerEarningScaled = nextPrice*PRECISION_BASE - result.feeScaled
	}

	// bookkeeping for multiple purchases
	t.amountLeft -= nextPrice
	t.paintedPixelCount++
	t.totalFeesScaled += result.feeScaled

	return result
}

func (t *TransferContext) hasReferrer() bool { return t.referrer != 0 }

func (t *TransferContext) updateFeesDistribution() {
	t.patronBonusScaled = t.totalFeesScaled * PATRON_BONUS_PERCENTAGE_POINTS / 100

	t.potScaled = t.totalFeesScaled * POT_PERCENTAGE_POINTS / 100

	t.referralEarningScaled = t.totalFeesScaled * REFERRER_PERCENTAGE_POINTS / 100
	if t.referrer == 0 {
		// if no referrer, pay the pot.
		t.potScaled += t.referralEarningScaled
		t.referralEarningScaled = 0
	}

	t.teamScaled = t.totalFeesScaled - t.patronBonusScaled - t.potScaled - t.referralEarningScaled
}

func (t *TransferContext) updateCanvas(cv *Canvas) {
	cv.potScaled += t.potScaled
	cv.teamScaled += t.teamScaled
	cv.pixelsDrawn += t.paintedPixelCount

	t.bonusPerPixelScaled = t.patronBonusScaled / cv.pixelsDrawn
	// eosio_assert(bonusPerPixelScaled > 0, "bonus is 0")

	cv.maskScaled += t.bonusPerPixelScaled
	// eosio_assert(cv.maskScaled >= bonusPerPixelScaled, "canvas mask overflow")

	t.canvasMaskScaled = cv.maskScaled
}

func (t *TransferContext) updatePurchaserAccount(acct *Account) {
	if t.amountLeft > 0 {
		acct.balanceScaled += t.amountLeft * PRECISION_BASE
	}

	patronBonusScaled := t.bonusPerPixelScaled * t.paintedPixelCount

	// FIXME: give the truncated to the team?

	maskUpdate := t.canvasMaskScaled*t.paintedPixelCount - patronBonusScaled
	acct.maskScaled += maskUpdate

	// eosio_assert(acct.maskScaled >= maskUpdate, "player mask overflow");
	acct.pixelsDrawn += t.paintedPixelCount
}

type Transfer struct {
}
