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

	fmt.Println("data:", userMap)
}
