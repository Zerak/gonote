package main

//import (
//	"fmt"
//	"math/big"
//)
//
//const (
//	// block per day
//	BLOCK_PER_DAY = 6000
//
//	MAX_REWARD_JOYS_PER_DAY   = 20000
//	MAX_REWARD_JOYS_PER_BLOCK = float64(MAX_REWARD_JOYS_PER_DAY) / float64(BLOCK_PER_DAY)
//
//	// BONUS_MULTIPLIER_8 Bonus muliplier for early joys makers.
//	//  1-12 8*64 = 512
//	// 13-24 4*64 = 256
//	// 25-36 2*64 = 128
//	// 37-48 1*64 = 64
//	DEFAULT_MULTIPLIER = 20
//	//BONUS_MULTIPLIER_8 = 20 * 10
//	//BONUS_MULTIPLIER_4 = 10 * 10
//	//BONUS_MULTIPLIER_2 = 5 * 10
//	//BONUS_MULTIPLIER_1 = 2.5 * 10
//	BONUS_MULTIPLIER_8 = 20
//	BONUS_MULTIPLIER_4 = 10
//	BONUS_MULTIPLIER_2 = 5
//	BONUS_MULTIPLIER_1 = 2.5
//
//	// mined 48 days
//	MINED_DAYS = 48
//
//	// continue days per every mine stage
//	CONTINUE_DAYS_PER_STAGE = 12
//
//	bonusBeginBlock = 100
//
//	DEFAULT_HEROES = 100
//)
//
//type PeriodInfo struct {
//	begin      int
//	end        int
//	multiplier int
//	times      float64
//}
//
//var (
//	lastMinersCount = DEFAULT_HEROES
//
//	periodInfo    []PeriodInfo
//	heroesInfo    = map[int]int{}
//	periodInfoOri []PeriodInfo
//)
//
//func getMultiplier(_from, _to int) int {
//	length := len(periodInfo)
//	totalMulti := 0
//	begin := 0
//	end := 0
//	multi := 0
//	for pid := 0; pid < length; pid++ {
//		info := periodInfo[pid]
//		begin = info.begin
//		end = info.end
//		multi = info.multiplier
//		if _to <= end {
//			if pid == 0 {
//				return (_to-_from)*multi + totalMulti
//			} else {
//				return (_to-begin)*multi + totalMulti
//			}
//		} else if _from >= end {
//			continue
//		} else {
//			if pid == 0 {
//				totalMulti = (end-_from)*multi + totalMulti
//			} else {
//				totalMulti = (end-begin)*multi + totalMulti
//			}
//		}
//	}
//	if _to > end {
//		totalMulti = totalMulti + (_to - end)
//	}
//
//	return totalMulti
//}
//
//func getReward(_from, _to int) *big.Float {
//	totalReward := big.NewFloat(0)
//	//fIdx := 0
//	//tIdx := 0
//	//for i := 0; i < len(periodInfo); i++ {
//	//	if periodInfo[i].begin <= _from {
//	//		fIdx = i
//	//	}
//	//	if periodInfo[i].end >= _to {
//	//		tIdx = i
//	//		break
//	//	}
//	//}
//	//if fIdx == tIdx {
//	//	totalReward.Add(totalReward, big.NewFloat(float64(_to-_from)*periodInfo[fIdx].times))
//	//} else {
//	//	for i := fIdx; i <= tIdx; i++ {
//	//		if i == fIdx {
//	//			totalReward.Add(totalReward, big.NewFloat(float64(periodInfo[i].end-_from)*periodInfo[i].times))
//	//		} else if i == tIdx {
//	//			totalReward.Add(totalReward, big.NewFloat(float64(_to-periodInfo[i].begin)*periodInfo[i].times))
//	//		} else {
//	//			totalReward.Add(totalReward, big.NewFloat(float64(periodInfo[i].end-periodInfo[i].begin)*periodInfo[i].times))
//	//		}
//	//	}
//	//}
//
//	// func2
//	for i := 0; i < len(periodInfo); i++ {
//		if _to <= periodInfo[i].end {
//			if i == 0 {
//				totalReward.Add(totalReward, big.NewFloat(float64(_to-_from)*periodInfo[i].times))
//				break
//			} else {
//				dest := periodInfo[i].begin
//				if _from > periodInfo[i].begin {
//					dest = _from
//				}
//				totalReward.Add(totalReward, big.NewFloat(float64(_to-dest)*periodInfo[i].times))
//				break
//			}
//		} else if _from >= periodInfo[i].end {
//			continue
//		} else {
//			totalReward.Add(totalReward, big.NewFloat(float64(periodInfo[i].end-_from)*periodInfo[i].times))
//			_from = periodInfo[i].end
//		}
//	}
//
//	// periodInfo multi times 10
//	// minersInfo multi times 1e12, 1e18/(10*1e12)
//	//return totalReward * 1e5
//	return totalReward.Mul(totalReward, big.NewFloat(1e18))
//}
//
//func getRate(_block int) int {
//	for i := 0; i < len(periodInfo); i++ {
//		if _block <= periodInfo[i].end {
//			return periodInfo[i].multiplier
//		}
//	}
//	return 0
//}
//
//func init() {
//	bonusEndBlock := bonusBeginBlock + BLOCK_PER_DAY*MINED_DAYS
//
//	var (
//		multiplier   = BONUS_MULTIPLIER_8
//		currentBlock = bonusBeginBlock
//		lastBlock    = currentBlock
//	)
//	for ; currentBlock < bonusEndBlock; currentBlock += CONTINUE_DAYS_PER_STAGE * BLOCK_PER_DAY {
//		periodInfo = append(periodInfo, PeriodInfo{
//			begin:      lastBlock,
//			end:        currentBlock + CONTINUE_DAYS_PER_STAGE*BLOCK_PER_DAY,
//			multiplier: multiplier,
//			times:      DEFAULT_HEROES * float64(multiplier) / BLOCK_PER_DAY,
//		})
//
//		//lastBlock = currentBlock + CONTINUE_DAYS_PER_STAGE*BLOCK_PER_DAY + 1
//		lastBlock = currentBlock + CONTINUE_DAYS_PER_STAGE*BLOCK_PER_DAY
//		multiplier = multiplier / 2
//	}
//
//	periodInfoOri = append(periodInfoOri, periodInfo...)
//}
//
//func totalHeroes() int {
//	return len(heroesInfo)
//}
//
//func updatePeriod(blockNumber int) {
//	if len(heroesInfo) < lastMinersCount+DEFAULT_HEROES {
//		return
//	}
//	// update last miners count
//	lastMinersCount = len(heroesInfo)
//
//	// push last item
//	lastItem := PeriodInfo{
//		begin:      blockNumber,
//		end:        0,
//		multiplier: 0,
//		times:      0,
//	}
//	periodInfo = append(periodInfo, lastItem)
//
//	for i := len(periodInfo) - 1; i > 0; i-- {
//		if periodInfo[i-1].begin < blockNumber {
//			end := periodInfo[i-1].end
//			multiplier := periodInfo[i-1].multiplier
//
//			//periodInfo[i-1].end = blockNumber - 1
//			periodInfo[i-1].end = blockNumber
//
//			times := float64(lastMinersCount) * float64(multiplier) / BLOCK_PER_DAY
//			if times >= MAX_REWARD_JOYS_PER_BLOCK {
//				times = MAX_REWARD_JOYS_PER_BLOCK
//			}
//			periodInfo[i].end = end
//			periodInfo[i].multiplier = multiplier
//			periodInfo[i].times = times
//			break
//		}
//
//		tmp := periodInfo[i-1]
//		tmp.times = float64(lastMinersCount) * float64(periodInfo[i-1].multiplier) / BLOCK_PER_DAY
//
//		periodInfo[i-1] = periodInfo[i]
//		periodInfo[i] = tmp
//	}
//}
//
//func main() {
//	for i := 0; i < 1500; i += 2 {
//		// 增加hero，获取奖励
//		heroesInfo[i] = i
//
//		// first update
//		updatePeriod(bonusBeginBlock + i)
//	}
//
//	reward := func(f, t int) {
//		reward := getReward(f, t)
//		reward.Quo(reward, big.NewFloat(1000000000000000000)) // div 1e18
//		fmt.Println("from:", f, " to:", t, " reward:", reward)
//	}
//
//	var (
//		v1 = 0.3333333333333333
//		v2 = 0.6666666666666666
//		v3 = 1.0
//		v4 = 1.3333333333333333
//		v5 = 1.6666666666666667
//		v6 = 2.0
//		v7 = 2.3333333333333335
//		//v8 = 0.16666666666666666
//	)
//	from := 100
//	to := 101
//	reward(from, to)
//	fmt.Println("wright:", v1*(101-100))
//
//	from = 101
//	to = 480
//	reward(from, to) // (480 - 101) * 0.333
//	fmt.Println("wright:", v1*(480-101))
//
//	from = 101
//	to = 650
//	reward(from, to) // (498-101) * 0.333) + (650 - 498) * 0.666
//	fmt.Println("wright:", v1*(498-101)+(650-498)*v2)
//
//	from = 101
//	to = 880
//	reward(from, to)
//	fmt.Println("wright:", v1*(498-101)+(698-498)*v2+(880-698)*v3)
//
//	from = 101
//	to = 1000
//	reward(from, to)
//	fmt.Println("wright:", v1*(498-101)+(698-498)*v2+(898-698)*v3+(1000-898)*v4)
//
//	from = 101
//	to = 1100
//	reward(from, to)
//	fmt.Println("wright:", v1*(498-101)+(698-498)*v2+(898-698)*v3+(1098-898)*v4+(1100-1098)*v5)
//
//	// 下一区块
//	from = 580
//	to = 680
//	reward(from, to)
//	fmt.Println("wright:", (680-580)*v2)
//
//	from = 580
//	to = 880
//	reward(from, to)
//	fmt.Println("wright:", (698-580)*v2+(880-698)*v3)
//
//	from = 580
//	to = 1000
//	reward(from, to)
//	fmt.Println("wright:", (698-580)*v2+(898-698)*v3+(1000-898)*v4)
//
//	from = 580
//	to = 1100
//	reward(from, to)
//	fmt.Println("wright:", (698-580)*v2+(898-698)*v3+(1098-898)*v4+(1100-1098)*v5)
//
//	// 下下个区块跨区块
//	from = 1100
//	to = 70000
//	reward(from, to)
//	fmt.Println("wright:", (1297-1100)*v5+(1498-1298)*v6+(70000-1498)*v7)
//}
