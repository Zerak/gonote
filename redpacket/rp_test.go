package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestGetRandom(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		rp := RedPacket{
			RemainSize:  2,
			RemainMoney: 10,
		}
		m := getRandomMoney(rp)
		fmt.Println(m)
	}
}
