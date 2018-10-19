package main

// startIndex,endIndex-1
// len = endIndex - startIndex
// cap = len(arr) - startIndex
import "fmt"

// todo
func main() {
	arr := [6]int{1, 2, 3, 4, 5, 6}
	s1 := arr[0:4]
	fmt.Print("s1 len:", len(s1), " cap:", cap(s1))
	fmt.Println(s1)
	s11 := arr[0:]
	fmt.Print("s11 len:", len(s11), " cap:", cap(s11))
	fmt.Println(s11)

	s2 := s1[2:]
	fmt.Println("s2 len:", len(s2), " cap:", cap(s2))
	fmt.Println(s2)

	s3 := s1[3:]
	fmt.Println("s3 len:", len(s3), " cap:", cap(s3))
	fmt.Println(s3)
}
