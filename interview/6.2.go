package main

import "fmt"

func main() {
	arr1 := [5]int{1, 2, 3, 4, 5}
	slice1 := arr1[1:2]
	fmt.Println("slice1:", slice1)

	slice1 = append(slice1, 6, 7, 8)
	fmt.Println("slice1:", slice1)
	fmt.Println("arr1:", arr1)

	arr2 := [5]int{1, 2, 3, 4, 5}
	slice2 := arr2[1:3]
	slice2 = append(slice2, 6, 7, 8)
	fmt.Println("slice2:", slice2)
	fmt.Println("arr2:", arr2)
}
