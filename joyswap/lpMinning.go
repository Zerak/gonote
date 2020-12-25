package main

import (
	"fmt"
	"math/big"
	"math/rand"
)

const (
	// block per day
	BLOCK_PER_DAY = 6100

	// JOYS tokens mined per block.
	JOYS_PER_BLOCK = 64

	// BONUS_MULTIPLIER_8 Bonus muliplier for early joys makers.
	BONUS_MULTIPLIER_8 = 120
	BONUS_MULTIPLIER_4 = 60
	BONUS_MULTIPLIER_2 = 30
	BONUS_MULTIPLIER_1 = 15

	// mined 48 days
	MINED_DAYS = 48

	// continue days per every mine stage
	CONTINUE_DAYS_PER_STAGE = 12

	bonusBeginBlock = 100

	DEFAULT_WETH = 100
)

type PeriodInfo struct {
	begin      int
	end        int
	multiplier int
}

var (
	lastWETHAmount = DEFAULT_WETH

	periodInfo    []PeriodInfo
	ethCount      = 0
	periodInfoOri []PeriodInfo
)

func init() {
	bonusEndBlock := bonusBeginBlock + BLOCK_PER_DAY*MINED_DAYS

	var (
		multiplier   = BONUS_MULTIPLIER_8
		currentBlock = bonusBeginBlock
		lastBlock    = currentBlock
	)
	for ; currentBlock < bonusEndBlock; currentBlock += CONTINUE_DAYS_PER_STAGE * BLOCK_PER_DAY {
		periodInfo = append(periodInfo, PeriodInfo{
			begin:      lastBlock,
			end:        currentBlock + CONTINUE_DAYS_PER_STAGE*BLOCK_PER_DAY,
			multiplier: multiplier,
		})

		//lastBlock = currentBlock + CONTINUE_DAYS_PER_STAGE*BLOCK_PER_DAY + 1
		lastBlock = currentBlock + CONTINUE_DAYS_PER_STAGE*BLOCK_PER_DAY
		multiplier = multiplier / 2
	}

	periodInfoOri = append(periodInfoOri, periodInfo...)
}

func deposit() {
	for i := 0; i < rand.Intn(120)+1; i++ {
		ethCount += i
	}
}

func getReward(_from, _to int) *big.Float {
	totalReward := big.NewFloat(0)
	for i := 0; i < len(periodInfo); i++ {
		if _to <= periodInfo[i].end {
			if i == 0 {
				totalReward.Add(totalReward, big.NewFloat(float64((_to-_from)*periodInfo[i].multiplier)))
				break
			} else {
				dest := periodInfo[i].begin
				if _from > periodInfo[i].begin {
					dest = _from
				}
				totalReward.Add(totalReward, big.NewFloat(float64((_to-dest)*periodInfo[i].multiplier)))
				break
			}
		} else if _from >= periodInfo[i].end {
			continue
		} else {
			totalReward.Add(totalReward, big.NewFloat(float64((periodInfo[i].end-_from)*periodInfo[i].multiplier)))
			_from = periodInfo[i].end
		}
	}
	totalReward.Mul(totalReward, big.NewFloat(1e17))
	return totalReward
}

func updatePeriod() {
	if lastWETHAmount >= 1000 {
		return
	}

	if ethCount < lastWETHAmount+DEFAULT_WETH {
		return
	}
	delt := (ethCount - lastWETHAmount) / DEFAULT_WETH

	// update last eth amount
	lastWETHAmount = ethCount

	for i := 0; i < len(periodInfo); i++ {
		periodInfo[i].multiplier = periodInfo[i].multiplier + delt*60
	}
}

func main() {
	fmt.Println(periodInfo)

	for i := 0; i < 1500; i += 2 {
		// 增加eth
		deposit()

		// first update
		updatePeriod()
	}
	fmt.Println(periodInfo)

	reward := func(f, t int) {
		reward := getReward(f, t)
		reward.Quo(reward, big.NewFloat(100000000000000000)) // div 1e17
		fmt.Println("from:", f, " to:", t, " reward:", reward)
	}

	var (
		v1 = 600
		v2 = 540
		v3 = 510
		v4 = 495
	)
	from := 100
	to := 101
	reward(from, to)
	fmt.Println("wright:", v1*(101-100))

	from = 101
	to = 73305
	reward(from, to) // (480 - 101) * 0.333
	fmt.Println("wright:", v1*(73300-101)+v2*(73305-73300))

	// 下一区块
	from = 73308
	to = 73408
	reward(from, to)
	fmt.Println("wright:", (73408-73308)*v2)

	from = 73308
	to = 219608
	reward(from, to)
	fmt.Println("wright:", (146500-73308)*v2+(219608-146500)*v3)

	from = 73308
	to = 292808
	reward(from, to)
	fmt.Println("wright:", (146500-73308)*v2+(219700-146500)*v3+(292808-219700)*v4)
}
