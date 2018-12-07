package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	////getID := func() (uuid.UUID, error) {
	//_ := func() (uuid.UUID, error) {
	//	return uuid.NewUUID()
	//}

	//idMap := make(map[string]bool)
	for i := 0; i < 1000; i++ {
		// id, err := getID()
		id, err := uuid.NewUUID()
		if err != nil {
			panic(err)
		}
		//if _, ok := idMap[id.String()]; ok {
		//	panic(fmt.Errorf("key same:%s", id.String()))
		//}
		//idMap[id.String()] = true
		fmt.Println(id.String())
	}
}
