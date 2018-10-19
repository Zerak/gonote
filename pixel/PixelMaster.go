package main

import "time"

type PixelMaster struct {
	CanvasStore  map[int]*Canvas
	PixelStore   map[int]*PixelRow
	AccountStore map[int]*Account
	GuardStore   map[int]*Guard
}

func onTransfer(trans *Transfer) {

}

func init() {

}

func refresh() {

}

func changedur(duration time.Duration) {

}

func (p *PixelMaster) End() {
	if len(p.CanvasStore) <= 0 {
		return
	}

	// anyone can create new canvas
	canvas := p.CanvasStore[0]
	if !canvas.IsEnded() {
		return
	}

	if canvas.IsEnded() {
		p.CanvasStore = make(map[int]*Canvas)
	}

	p.CanvasStore[0].id = canvas.id + 1
	p.CanvasStore[0].lastPaintedAt = time.Now()
	p.CanvasStore[0].duration = CANVAS_DURATION
}

func createacct(account int) {

}

func withdraw(toAccount int) {

}

func (p *PixelMaster) Clearpixels(count, nonce int) {
	// todo count
	p.PixelStore = make(map[int]*PixelRow)
}

func (p *PixelMaster) Clearaccts(count, nonce int) {
	// todo count
	p.AccountStore = make(map[int]*Account)
}

func (p *PixelMaster) Clearcanvs(count, nonce int) {
	// todo count
	p.CanvasStore = make(map[int]*Canvas)
}

func (p *PixelMaster) Resetquota() {
	p.GuardStore = make(map[int]*Guard)
	/*
		guardItr = p.GuardStore.begin();
		if (guardItr == guards.end()) {
			guards.emplace(_self, [&](guard &grd) {
				grd.id = 0;
				grd.quota = WITHDRAW_QUOTA;
			});
		} else {
			guards.modify(guardItr, 0, [&](guard &grd) { grd.quota = WITHDRAW_QUOTA; });
		}
	*/
}

// func apply(account_name contract, action_name act){}
func apply(account_name, action_name string) {

}

// func isValidReferrer(account_name name)bool{
func isValidReferrer(name int) bool {
	return false
}

func (p *PixelMaster) deposit(user int, quantityScaled int) {
	// eosio_assert(quantityScaled > 0, "must deposit positive quantity");
	for k, v := range p.AccountStore {
		if v.owner == user {
			p.AccountStore[k].balanceScaled += quantityScaled
			break
		}
	}
}

func (p *PixelMaster) drawPixel(allPixels *map[int]*PixelRow, pixelOrder *PixelOrder, ctx *TransferContext) {
	loc := pixelOrder.Location()
	var pixel Pixel
	for _, v := range *allPixels {
		if v.row == loc.row {
			pixel = v.pixels[loc.col]
		}
	}

	// TODO extract this into its own method
	// Emplace & initialize empty row if it doesn't already exist

	result := ctx.purchase(&pixel, pixelOrder)
	if result.isSkipped {
		return
	}

	for _, v := range *allPixels {
		v.pixels[loc.col].color = pixelOrder.color
		v.pixels[loc.col].priceCounter = pixel.NextPriceCounter()
		v.pixels[loc.col].owner = ctx.purchaser
	}

	if !result.isFirstBuyer {
		p.deposit(pixel.owner, result.ownerEarningScaled)
	}
}

func refreshLastPaintedAt() {

}
