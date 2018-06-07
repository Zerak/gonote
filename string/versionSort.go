package main

import (
	"fmt"
	"sort"
)
import "strings"

//var unsorted = []string{"1.6beta1", "1.5rc1", "1.5beta2", "1.5beta1", "1.5.1", "1.5", "1.4rc2", "1.4rc1", "1.4beta1", "1.4.2", "1.4.1", "1.4", "1.3rc2", "1.3rc1", "1.3beta2", "1.3beta1", "1.3.3", "1.3.2", "1.3.1", "1.3", "1.2rc5", "1.2rc4", "1.2rc3", "1.2rc2", "1.2rc1", "1.2.2", "1.2.1", "1.2", "1.1.2", "1.1.1", "1.1", "1.0.3", "1.0.2", "1.5.2", "1.5alpha1"}
var unsorted = []string{"11.1.2", "12.2.3", "9.5.4"}

func main() {
	aaa()
	return
	sorted := VersionSort(unsorted)
	for _, v := range sorted {
		println(v)
	}
}

func VersionSort(versions []string) []string {
	s2 := make([]string, len(versions))
	for i, v := range versions {
		s2[i] = strings.Replace(v, ".", "~", -1) + "~"
	}
	sort.Strings(s2)
	for i, v := range s2 {
		s2[i] = strings.Replace(v[:len(v)-1], "~", ".", -1)
	}
	return s2
}

func aaa() {
	recordUid := []int{
		3749,
		3733,
		3665,
		3662,
		3650,
		3641,
		3635,
		3633,
		3613,
		3601,
		3593,
		3572,
		3568,
		3564,
		3527,
		3504,
		3497,
		3495,
		3484,
		3442,
		3439,
		3438,
		3432,
		3421,
		3415,
		3412,
		3393,
		3388,
		3372,
		3351,
		3350,
		3327,
		3322,
		3280,
		3235,
		3207,
		3134,
		3102,
		3070,
		2926,
		2908,
		2897,
		2810,
		2799,
		2673,
		2644,
		2631,
		2594,
		2549,
		2459,
		2414,
		2396,
		2381,
		2334,
		2291,
		2233,
		2137,
		2117,
		2092,
		2009,
		1999,
		1980,
		1971,
		1919,
		1910,
		1887,
		1832,
		1813,
		1779,
		1666,
		1655,
		1651,
		1649,
		1636,
		1609,
		1528,
		1282,
		1211,
		1136,
		1111,
		1070,
		1051,
		1045,
		1030,
		982,
		971,
		944,
		929,
		925,
		913,
		897,
		843,
		826,
		794,
		757,
		751,
		750,
		738,
		734,
		720,
		327,
		202,
		45,
	}

	rewardUid := []int{
		3432,
		3393,
		3327,
		3322,
		3280,
		3102,
		2897,
		2631,
		2549,
		2414,
		2137,
		2092,
		2009,
		1919,
		1832,
		1813,
		1666,
		1651,
		1649,
		1211,
		1136,
		1111,
		982,
		971,
		929,
		897,
		843,
		826,
		734,
		720,
		327,
		45,
	}

	size := 0
	for _, v := range recordUid {
		find := false
		for i := 0; i < len(rewardUid); i++ {
			if rewardUid[i] == v {
				find = true
				break
			}
		}
		if !find {
			size++
			fmt.Println(v)
		}
	}
	fmt.Println("size:", size)
}
