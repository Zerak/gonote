package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type RedPacket struct {
	RemainSize  int     // 剩余的红包数量
	RemainMoney float64 // 剩余的钱
}

func getRandomMoney(rp RedPacket) float64 {
	if rp.RemainSize == 1 {
		rp.RemainSize--
		return math.Ceil(rp.RemainMoney*100) / 100
	}

	var (
		money float64
		min   float64
		max   float64
	)

	min = 0.01
	max = rp.RemainMoney / float64(rp.RemainSize) * 2
	money = rand.Float64() * max
	if money <= 0.01 {
		money = min
	}

	return math.Floor(money*100) / 100
}

func main() {
	rand.Seed(time.Now().UnixNano())
	rp := RedPacket{
		RemainSize:  10,
		RemainMoney: 100,
	}
	m := getRandomMoney(rp)
	fmt.Println(m)
}
