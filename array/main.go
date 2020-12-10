package main

import (
	"fmt"
	"math/rand"
	"time"
)

func t() {
}

func main() {
	tick := 0

	//level := sort.IntSlice{}
	level := make([]int, 6)
	level[0] = 10000
	level[1] = 5000
	level[2] = 2000
	level[3] = 500
	level[4] = 100
	level[5] = 5

	//level[0] = 10000
	//level[1] = 4000
	//level[2] = 2000
	//level[3] = 500
	//level[4] = 0
	//level[5] = 0

	add := func(step int) int {
		sum := 0
		idx := 0
		for i := len(level); i > 0; i-- {
			idx++
			sum += level[i-1]
			if idx >= step {
				break
			}
		}
		return sum
	}
	////ll := [7]int{0, 5, 105, 605, 2605, 7506, 17605}
	//llTmp := [7]int{}
	//llTmp[0] = 0
	//for i := len(level); i > 0; i-- {
	//	llTmp[len(level)-i+1] = add(len(level) - i + 1)
	//}

	//ll := [7]int{17605, 7506, 2605, 605, 105, 5, 0}
	ll := [7]int{}
	for i := 0; i < len(level); i++ {
		ll[i] = add(len(level) - i)
	}

	//idx := 0
	//for i := len(llTmp); i > 0; i-- {
	//	ll[idx] = llTmp[i-1]
	//	idx++
	//}
	parseLevel := func(weight int) int {
		var (
			k int
			l int
			v int
		)
		for k, v = range ll {
			l = k
			if weight >= v {
				if k == 0 {
					//l = k + 1
					return 1
				}
				break
			}
			if weight < v {
				continue
			}
		}
		if tick <= 200 {
			if l == 6 {
				t()
				fmt.Println(tick, "次,抽取6级女神")
			}
			if l == 5 {
				t()
				fmt.Println(tick, "次,抽取5级索尔")
			}
			//if l == 4 {
			//	t()
			//	fmt.Println(tick, "次,抽取4级星际猎人")
			//}
			//if l == 3 {
			//	t()
			//	fmt.Println(tick, "次,抽取3级机甲")
			//}
		}

		return l
	}

	gas := 0
	for _, v := range level {
		gas += v
	}
	total := 100000
	sum := map[int]int{}

	for tick = 0; tick < total; tick++ {
		rand.Seed(time.Now().UnixNano())
		weight := rand.Intn(gas + 1)
		l := parseLevel(weight)
		sum[l]++
	}
	fmt.Println("sum:", sum)

	for k, v := range sum {
		fmt.Println("level:", k, " percent:", float64(v)/float64(total)*100.0, "%")
	}
}
