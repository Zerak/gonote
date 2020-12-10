package main

//import "fmt"
//
//const (
//	// block per day
//	BLOCK_PER_DAY = 6100
//
//	// JOYS tokens mined per block.
//	JOYS_PER_BLOCK = 64
//
//	// BONUS_MULTIPLIER_8 Bonus muliplier for early joys makers.
//	//  1-12 8*64 = 512
//	// 13-24 4*64 = 256
//	// 25-36 2*64 = 128
//	// 37-48 1*64 = 64
//	BONUS_MULTIPLIER_8 = 8
//	BONUS_MULTIPLIER_4 = 4
//	BONUS_MULTIPLIER_2 = 2
//	BONUS_MULTIPLIER_1 = 1
//
//	// mined 48 days
//	MINED_DAYS = 48
//
//	// continue days per every mine stage
//	CONTINUE_DAYS_PER_STAGE = 12
//
//	bonusBeginBlock = 100
//)
//
//type PeriodInfo struct {
//	begin      int
//	end        int
//	multiplier int
//}
//
//var periodInfo []PeriodInfo
//
//func main() {
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
//		})
//
//		lastBlock = currentBlock + CONTINUE_DAYS_PER_STAGE*BLOCK_PER_DAY + 1
//		multiplier = multiplier / 2
//	}
//
//	fmt.Println(periodInfo)
//
//	// 100-73300 73301-146500 146501-219700 219701-292900
//	testCase0()
//	fmt.Println()
//
//	testCase1()
//	fmt.Println()
//
//	testCase2()
//	fmt.Println()
//
//	testCase3()
//	fmt.Println()
//
//	testCase4()
//	fmt.Println()
//
//	testCase5()
//	fmt.Println()
//}
//
//func testCase0() {
//	fmt.Println("testCase0")
//	// 100-73300 73301-146500 146501-219700 219701-292900
//	multi := getMultiplier(90, 101)
//	fmt.Println(multi, multi == (101-100+1)*BONUS_MULTIPLIER_8)
//
//	multi = getMultiplier(90, 73305)
//	fmt.Println(multi, multi == (73300-100+1)*BONUS_MULTIPLIER_8+(73305-73301+1)*BONUS_MULTIPLIER_4)
//}
//
//func testCase1() {
//	fmt.Println("testCase1")
//	// 100-73300 73301-146500 146501-219700 219701-292900
//	multi := getMultiplier(100, 101)
//	fmt.Println(multi, multi == (101-100+1)*BONUS_MULTIPLIER_8)
//
//	multi = getMultiplier(100, 73305)
//	fmt.Println(multi, multi == (73300-100+1)*BONUS_MULTIPLIER_8+(73305-73301+1)*BONUS_MULTIPLIER_4)
//
//	multi = getMultiplier(100, 146503)
//	fmt.Println(multi, multi == (73300-100+1)*BONUS_MULTIPLIER_8+(146500-73301+1)*BONUS_MULTIPLIER_4+(146503-146501+1)*BONUS_MULTIPLIER_2)
//
//	multi = getMultiplier(100, 219703)
//	fmt.Println(multi, multi == (73300-100+1)*BONUS_MULTIPLIER_8+(146500-73301)*BONUS_MULTIPLIER_4+(219700-146501+1)*BONUS_MULTIPLIER_2+(219703-219701+1)*BONUS_MULTIPLIER_1)
//
//	multi = getMultiplier(100, 292903)
//	fmt.Println(multi, multi == (73300-100+1)*BONUS_MULTIPLIER_8+(146500-73301+1)*BONUS_MULTIPLIER_4+(219700-146501+1)*BONUS_MULTIPLIER_2+(292900-219701+1)*BONUS_MULTIPLIER_1)
//}
//
//func testCase2() {
//	fmt.Println("testCase2")
//	// 100-73300 73301-146500 146501-219700 219701-292900
//	multi := getMultiplier(73305, 73309)
//	fmt.Println(multi, multi == (73309-73305+1)*BONUS_MULTIPLIER_4)
//
//	multi = getMultiplier(73305, 146503)
//	fmt.Println(multi, multi == (146500-73305+1)*BONUS_MULTIPLIER_4+(146503-146501+1)*BONUS_MULTIPLIER_2)
//
//	multi = getMultiplier(73305, 219705)
//	fmt.Println(multi, multi == (146500-73305+1)*BONUS_MULTIPLIER_4+(219700-146501+1)*BONUS_MULTIPLIER_2+(219705-219701+1)*BONUS_MULTIPLIER_1)
//
//	multi = getMultiplier(73300, 73309)
//	fmt.Println(multi, multi == (73309-73305+1)*BONUS_MULTIPLIER_4)
//}
//
//func testCase3() {
//	fmt.Println("testCase3")
//	// 100-73300 73301-146500 146501-219700 219701-292900
//}
//
//func testCase4() {
//	fmt.Println("testCase4")
//	// 100-73300 73301-146500 146501-219700 219701-292900
//}
//
//func testCase5() {
//	fmt.Println("testCase5")
//	// 100-73300 73301-146500 146501-219700 219701-292900
//}
//
//// Return reward multiplier over the given _from to _to block.
//func getMultiplier(_from, _to int) int {
//	length := len(periodInfo)
//	var (
//		begin, end, multi, totalMulti int
//	)
//
//	if _from < bonusBeginBlock {
//		_from = bonusBeginBlock
//	}
//
//	for pid := 0; pid < length; pid++ {
//		info := periodInfo[pid]
//		begin = info.begin
//		end = info.end
//		multi = info.multiplier
//
//		if _to <= end {
//			if pid == 0 {
//				return (_to-_from)*multi + totalMulti
//			} else {
//				if _from > begin {
//					begin = _from
//				}
//				return (_to-begin)*multi + totalMulti
//			}
//		} else if _from >= end {
//			continue
//		} else {
//			if pid == 0 {
//				totalMulti = (end-_from)*multi + totalMulti
//			} else {
//				if _from > begin {
//					begin = _from
//				}
//				totalMulti = (end-begin)*multi + totalMulti
//			}
//		}
//	}
//
//	return totalMulti
//}
