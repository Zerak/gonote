package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	n1 := 0
	n2 := 0
	n3 := 0
	n4 := 0

	count := 500000
	sum := 0
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < count; i++ {
		n := getProbabbity()
		if n >= 2 {
			if n == 2 {
				n2 += 1
			}
			if n == 3 {
				n3 += 1
			}
			if n == 4 {
				n4 += 1
			}
		}
		sum += getWithral(n)
	}
	n1 = count - n2 - n3 - n4
	fmt.Printf("grade->n1:%v n2:%v n3:%v n4:%v sum:%v\n", n1, n2, n3, n4, sum)
}

func getProbabbity() int {
	re := rand.Intn(1000000)
	if re > 5000 {
		return 1
	} else if re > 500 && re <= 5000 {
		return 2
	} else if re > 25 && re <= 500 {
		return 3
	} else if re <= 25 {
		fmt.Printf("re:%v\n", re)
		return 4
	}
	return 0
}

// 1,99.45
// 2,0.5
// 3,0.05
// 4,0.0025
func getProbabbity1() int {
	re := rand.Intn(10000)
	if re > 5000 {
		return 1
	} else if re > 475 && re <= 5000 {
		return 2
	} else if re > 25 && re <= 475 {
		return 3
	} else if re <= 25 {
		fmt.Printf("re:%v\n", re)
		return 4
	}
	return 0
}
func getProbabbity2() int {
	re := rand.Intn(100)
	if re > 5000 {
		return 1
	} else if re > 500 && re <= 5000 {
		return 2
	} else if re > 25 && re <= 500 {
		return 3
	} else if re <= 25 {
		fmt.Printf("re:%v\n", re)
		return 4
	}
	return 0
}

func getWithral(grade int) int {
	if grade == 1 {
		return rand.Intn(20) + 10 //  [10 ,30)
	}
	if grade == 2 {
		return rand.Intn(400) + 300 // [300,700)
	}
	if grade == 3 {
		return rand.Intn(400) + 800 // [800,1200)
	}
	if grade == 4 {
		return 88888
	}
	return rand.Intn(20) + 10 //  [10 ,30
}
