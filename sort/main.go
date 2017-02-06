package main

import (
	"fmt"
	"sort"
)

type Person struct {
	//Name string
	Name int
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v: %v", p.Name, p.Age)
}

// ByAge implements sort.Interface for []Person based on
// the Age field.
type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Name < a[j].Name }

func main() {
	//people := []Person{
	//	{"1", 31},
	//	{"11", 42},
	//	{"13", 17},
	//	{"5", 26},
	//}
	people := []Person{
		{1, 31},
		{11, 42},
		{13, 17},
		{5, 26},
	}

	fmt.Println(people)
	sort.Sort(ByAge(people))
	fmt.Println(people)

	if "1" > "11" {
		fmt.Printf("err\n")
	} else {
		fmt.Printf("ok\n")
	}
}
