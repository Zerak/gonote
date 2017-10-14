package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"sort"
	"time"
)

var tiles = []byte{
	0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, // Dots
	0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, // Bamboo
	0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, // Characters
	0x31, 0x41, 0x51, 0x61, 0x71, 0x81, 0x91, // East South West North Red Green White
}

var full = []byte{
	0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, // Dots
	0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, // Bamboo
	0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, // Characters
	0x31, 0x41, 0x51, 0x61, 0x71, 0x81, 0x91, // East South West North Red Green White
	0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, // Dots
	0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, // Bamboo
	0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, // Characters
	0x31, 0x41, 0x51, 0x61, 0x71, 0x81, 0x91, // East South West North Red Green White
	0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, // Dots
	0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, // Bamboo
	0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, // Characters
	0x31, 0x41, 0x51, 0x61, 0x71, 0x81, 0x91, // East South West North Red Green White
	0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, // Dots
	0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, // Bamboo
	0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, // Characters
	0x31, 0x41, 0x51, 0x61, 0x71, 0x81, 0x91, // East South West North Red Green White
}

func findPairs() [][]byte {
	pairs := make([][]byte, 0, len(tiles))

	for _, v := range tiles {
		pair := []byte{v, v}
		pairs = append(pairs, pair)
	}

	return pairs
}

func findGroups() [][]byte {
	groups := make([][]byte, 0, len(tiles)+(9-2)*3)

	// find three identical tiles
	for _, v := range tiles {
		group := []byte{v, v, v}
		groups = append(groups, group)
	}

	// find three sequence tiles
	for i := 2; i < len(tiles); i++ {
		if tiles[i-2]+1 == tiles[i-1] && tiles[i-1] == tiles[i]-1 {
			group := []byte{tiles[i-2], tiles[i-1], tiles[i]}
			groups = append(groups, group)
		}
	}

	return groups
}

type byteSlice []byte

func (b byteSlice) Len() int {
	return len(b)
}

func (b byteSlice) Less(i, j int) bool {
	return b[i] < b[j]
}

func (b byteSlice) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func checkValid(win []byte) bool {
	sort.Sort(byteSlice(win))

	for i := 4; i < len(win); i++ {
		if win[i] == win[i-4] {
			return false
		}
	}

	return true
}

func notExist(win []byte, wins [][]byte) bool {
	for _, v := range wins {
		if bytes.Equal(win, v) {
			return false
		}
	}

	return true
}

func composeWin(pairs, groups [][]byte) [][]byte {
	wins := make([][]byte, 0, 11498658)

	tmp := make([]byte, 14)
	for _, pair := range pairs {
		for _, group1 := range groups {
			for _, group2 := range groups {
				for _, group3 := range groups {
					for _, group4 := range groups {
						copy(tmp, pair)
						copy(tmp[2:], group1)
						copy(tmp[5:], group2)
						copy(tmp[8:], group3)
						copy(tmp[11:], group4)

						if checkValid(tmp) && notExist(tmp, wins) {
							win := make([]byte, 0, 14)
							win = append(win, tmp...)
							wins = append(wins, win)
						}
					}
				}
			}
		}
	}

	return wins
}

type twoUint64 struct {
	H uint64 // High
	L uint64 // Low
}

func composeWinEx(pairs, groups [][]byte) map[twoUint64][]byte {
	wins := make(map[twoUint64][]byte)

	var key twoUint64
	tmp := make([]byte, 14)
	for _, pair := range pairs {
		for _, group1 := range groups {
			for _, group2 := range groups {
				for _, group3 := range groups {
					for _, group4 := range groups {
						copy(tmp, pair)
						copy(tmp[2:], group1)
						copy(tmp[5:], group2)
						copy(tmp[8:], group3)
						copy(tmp[11:], group4)

						if checkValid(tmp) {
							key.H = uint64(tmp[0])
							key.L = uint64(tmp[6])

							for _, v := range tmp[1:6] {
								key.H = key.H<<8 + uint64(v)
							}

							for _, v := range tmp[7:] {
								key.L = key.L<<8 + uint64(v)
							}

							if _, ok := wins[key]; !ok {
								win := make([]byte, 0, 14)
								win = append(win, tmp...)
								wins[key] = win
							}
						}
					}
				}
			}
		}
	}

	return wins
}

type jsonData struct {
	H uint64 // High
	L uint64 // Low
	V []byte // Value
}

func toJSON() {
	pairs := findPairs()

	groups := findGroups()

	wins := composeWinEx(pairs, groups)

	f, err := os.Create("json.data")
	defer f.Close()

	if err != nil {
		log.Fatal("Create", err)
	}

	var jd jsonData

	enc := json.NewEncoder(f)

	for k, v := range wins {
		jd.H = k.H
		jd.L = k.L
		jd.V = v

		if err := enc.Encode(jd); err != nil {
			log.Fatal("Encode", err)
		}
	}
}

func benchmarkWin(n int, wins map[twoUint64][]byte) {
	var win int
	var key twoUint64
	now := time.Now()
	for i := 0; i < n; i++ {
		perm := rand.Perm(136)
		hand := make([]byte, 14)

		for j := 0; j < 14; j++ {
			hand[j] = full[perm[j]]
		}

		sort.Sort(byteSlice(hand))

		key.H = uint64(hand[0])
		key.L = uint64(hand[6])

		for _, v := range hand[1:6] {
			key.H = key.H<<8 + uint64(v)
		}

		for _, v := range hand[7:] {
			key.L = key.L<<8 + uint64(v)
		}

		if _, ok := wins[key]; ok {
			win++
		}
	}

	fmt.Printf("Test total %d, Win %d, Time %v\n", n, win, time.Since(now))
}

func benchmarkWinEx(n int, wins map[twoUint64][]byte) {
	var win int
	var key twoUint64
	now := time.Now()
	for i := 0; i < n; i++ {
		perm := rand.Perm(136)
		hand := make([]byte, 14)

		for j := 0; j < 14; j++ {
			hand[j] = full[perm[j]]
		}

	EXIT:
		for _, v1 := range tiles {
			for _, v2 := range tiles {
				for _, v3 := range tiles {
					tmp := make([]byte, 0, 14)
					tmp = append(tmp, v1, v2, v3)
					tmp = append(tmp, hand[3:]...)

					if checkValid(tmp) {
						key.H = uint64(tmp[0])
						key.L = uint64(tmp[6])

						for _, v := range tmp[1:6] {
							key.H = key.H<<8 + uint64(v)
						}

						for _, v := range tmp[7:] {
							key.L = key.L<<8 + uint64(v)
						}

						if _, ok := wins[key]; ok {
							win++
							break EXIT
						}
					}
				}
			}
		}
	}

	fmt.Printf("Test total %d, Win %d, Time %v\n", n, win, time.Since(now))
}

type simpleData struct {
	K int    // Key
	V []byte // Value
}

func bytesToInt(win []byte) int {
	tmp := make([]byte, 0, 17)
	tmp = append(tmp, 1)
	for i, pos := 1, 0; i < len(win); i++ {
		if win[i-1] == win[i] {
			tmp[pos]++
		} else if win[i-1]+1 == win[i] {
			tmp = append(tmp, 1)
			pos++
		} else {
			tmp = append(tmp, 1)
			tmp[pos] += 0x0A
			pos++
		}
	}

	res := 1
	for _, v := range tmp {
		switch v {
		case 0x01:
			res <<= 1
		case 0x02:
			res <<= 3
			res |= 0x06
		case 0x03:
			res <<= 5
			res |= 0x1E
		case 0x04:
			res <<= 7
			res |= 0x7E
		case 0x0B:
			res <<= 2
			res |= 0x02
		case 0x0C:
			res <<= 4
			res |= 0x0E
		case 0x0D:
			res <<= 6
			res |= 0x3E
		case 0x0E:
			res <<= 8
			res |= 0xFE
		}
	}

	return res
}

func toSimple(wins map[twoUint64][]byte) {
	f, err := os.Create("simple.data")
	defer f.Close()

	if err != nil {
		log.Fatal("Create", err)
	}

	var sd simpleData

	enc := json.NewEncoder(f)

	for _, win := range wins {
		sd.K = bytesToInt(win)
		sd.V = win

		if err := enc.Encode(sd); err != nil {
			log.Fatal("Encode", err)
		}
	}
}

func fromSimple() {
	f, err := os.Open("simple.data")
	defer f.Close()

	if err != nil {
		log.Fatal("Open", err)
	}

	var sd simpleData

	dec := json.NewDecoder(f)

	wins := make(map[int]bool)

	for {
		if err := dec.Decode(&sd); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal("Decode", err)
		}

		wins[sd.K] = true
	}

	benchmarkWin2(10000000, wins)
	benchmarkWinEx2(1000, wins)
}

func benchmarkWin2(n int, wins map[int]bool) {
	var win int
	now := time.Now()
	for i := 0; i < n; i++ {
		perm := rand.Perm(136)
		hand := make([]byte, 14)

		for j := 0; j < 14; j++ {
			hand[j] = full[perm[j]]
		}

		sort.Sort(byteSlice(hand))

		if _, ok := wins[bytesToInt(hand)]; ok {
			win++
		}
	}

	fmt.Printf("Test total %d, Win %d, Time %v\n", n, win, time.Since(now))
}

func benchmarkWinEx2(n int, wins map[int]bool) {
	var win int
	now := time.Now()
	for i := 0; i < n; i++ {
		perm := rand.Perm(136)
		hand := make([]byte, 14)

		for j := 0; j < 14; j++ {
			hand[j] = full[perm[j]]
		}

	EXIT:
		for _, v1 := range tiles {
			for _, v2 := range tiles {
				for _, v3 := range tiles {
					tmp := make([]byte, 0, 14)
					tmp = append(tmp, v1, v2, v3)
					tmp = append(tmp, hand[3:]...)

					if checkValid(tmp) {
						if _, ok := wins[bytesToInt(tmp)]; ok {
							win++
							break EXIT
						}
					}
				}
			}
		}
	}

	fmt.Printf("Test total %d, Win %d, Time %v\n", n, win, time.Since(now))
}

func benchmarkWin3(n int, wins map[twoUint64][]byte) {
	winsCopy := make(map[string][]byte)
	for _, v := range wins {
		winsCopy[string(v)] = v
	}

	var win int
	now := time.Now()
	for i := 0; i < n; i++ {
		perm := rand.Perm(136)
		hand := make([]byte, 14)

		for j := 0; j < 14; j++ {
			hand[j] = full[perm[j]]
		}

		sort.Sort(byteSlice(hand))

		if _, ok := winsCopy[string(hand)]; ok {
			win++
		}
	}

	fmt.Printf("Test total %d, Win %d, Time %v\n", n, win, time.Since(now))

	benchmarkWinEx3(1000, winsCopy)
}

func benchmarkWinEx3(n int, wins map[string][]byte) {
	var win int
	now := time.Now()
	for i := 0; i < n; i++ {
		perm := rand.Perm(136)
		hand := make([]byte, 14)

		for j := 0; j < 14; j++ {
			hand[j] = full[perm[j]]
		}

	EXIT:
		for _, v1 := range tiles {
			for _, v2 := range tiles {
				for _, v3 := range tiles {
					tmp := make([]byte, 0, 14)
					tmp = append(tmp, v1, v2, v3)
					tmp = append(tmp, hand[3:]...)

					if checkValid(tmp) {
						if _, ok := wins[string(tmp)]; ok {
							win++
							break EXIT
						}
					}
				}
			}
		}
	}

	fmt.Printf("Test total %d, Win %d, Time %v\n", n, win, time.Since(now))
}

func availableTiles(win []byte) map[byte]bool {
	available := make(map[byte]bool)

	for _, v := range win {
		if v > 0x01 && v < 0x09 || v > 0x11 && v < 0x19 || v > 0x21 && v < 0x29 {
			available[v-1], available[v+1] = true, true
		}

		available[v] = true
	}

	return available
}

func benchmarkWinEx3Ex(n int, wins map[string][]byte) {
	var win int
	now := time.Now()
	for i := 0; i < n; i++ {
		perm := rand.Perm(136)
		hand := make([]byte, 14)

		for j := 0; j < 14; j++ {
			hand[j] = full[perm[j]]
		}

	EXIT:
		for v1 := range availableTiles(hand[3:]) {
			hand[2] = v1
			for v2 := range availableTiles(hand[2:]) {
				hand[1] = v2
				for v3 := range availableTiles(hand[1:]) {
					hand[0] = v3

					tmp := make([]byte, 0, 14)
					tmp = append(tmp, hand...)

					if checkValid(tmp) {
						if _, ok := wins[string(tmp)]; ok {
							win++
							break EXIT
						}
					}
				}
			}
		}
	}

	fmt.Printf("Test total %d, Win %d, Time %v\n", n, win, time.Since(now))
}

func appendAvailableTiles(origin map[byte]int, ap ...byte) map[byte]int {
	available := make(map[byte]int)

	for k, v := range origin {
		available[k] = v
	}

	for _, v := range ap {
		if v > 0x01 && v < 0x09 || v > 0x11 && v < 0x19 || v > 0x21 && v < 0x29 {
			if _, ok := available[v-1]; !ok {
				available[v-1] = 0
			}

			if _, ok := available[v+1]; !ok {
				available[v+1] = 0
			}
		}

		available[v]++
	}

	return available
}

func benchmarkWinEx3Ex2(n int, wins map[string][]byte) {
	var win int
	now := time.Now()
	for i := 0; i < n; i++ {
		perm := rand.Perm(136)
		hand := make([]byte, 14)

		for j := 0; j < 14; j++ {
			hand[j] = full[perm[j]]
		}

		available1 := appendAvailableTiles(nil, hand[3:]...)

	EXIT:
		for v1 := range available1 {
			if available1[v1] >= 4 {
				continue
			}
			available2 := appendAvailableTiles(available1, v1)
			for v2 := range available2 {
				if available2[v2] >= 4 {
					continue
				}
				available3 := appendAvailableTiles(available2, v2)
				for v3 := range available3 {
					if available3[v3] >= 4 {
						continue
					}
					tmp := make([]byte, 0, 14)
					tmp = append(tmp, v1, v2, v3)
					tmp = append(tmp, hand[3:]...)

					sort.Sort(byteSlice(tmp))

					if _, ok := wins[string(tmp)]; ok {
						win++
						break EXIT
					}
				}
			}
		}
	}

	fmt.Printf("Test total %d, Win %d, Time %v\n", n, win, time.Since(now))
}

func main() {
	// fromSimple()
	// return
	f, err := os.Open("json.data")
	defer f.Close()

	if err != nil {
		log.Fatal("Open", err)
	}

	var jd jsonData

	dec := json.NewDecoder(f)

	wins := make(map[twoUint64][]byte)

	for {
		if err := dec.Decode(&jd); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal("Decode", err)
		}

		wins[twoUint64{
			H: jd.H,
			L: jd.L,
		}] = jd.V
	}

	benchmarkWin(10000000, wins)
	benchmarkWinEx(1000, wins)
	//benchmarkWin3(10000000, wins)
}
