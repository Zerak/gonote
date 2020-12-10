package main

import "fmt"

// todo
type User struct {
	UserID int
	Name   string
	Age    int
}

func main() {
	userMap := make(map[int]User)
	userMap[1] = User{UserID: 1, Name: "a", Age: 10}

	for _, v := range userMap {
		v.Name = "aNew"
		v.Age = 11
	}

	v := userMap[1]

	v.Name = "11"
	v.Age = 21

	fmt.Println("data:", userMap)
}
