package main

import (
	"fmt"
)

func binarySearch2(arr []int, want int) int {
	mid := 0
	start := 0
	end := len(arr) - 1
	for start <= end {
		mid = start + (end-start)/2
		if want < arr[mid] {
			end = mid - 1
		} else if want > arr[mid] {
			start = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func binarySearch(arr []int, start, end, want int) int {
	if start > end {
		return -1
	}

	mid := start + (end-start)/2
	if arr[mid] > want {
		return binarySearch(arr, start, mid-1, want)
	}
	if arr[mid] < want {
		return binarySearch(arr, mid+1, end, want)
	}
	return mid
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	fmt.Println(len(arr))

	for i := 0; i < 20; i++ {
		idx := binarySearch2(arr, i)
		fmt.Println("i:", i, " idx:", idx)
	}

	for i := 0; i < 20; i++ {
		idx := binarySearch(arr, 0, len(arr)-1, i)
		fmt.Println("i:", i, " idx:", idx)
	}
}
