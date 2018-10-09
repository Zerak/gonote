package main

import (
	"fmt"
	"gonote/factory/factory"
)

func main() {
	//datastore1, err := factory.GetDataStore("postgres")
	//datastore2, err := factory.GetDataStore("memory")

	datastore1, err := factory.GetDataStore(100)
	datastore2, err := factory.GetDataStore(9999)

	fmt.Println("d1:", datastore1.Name(), " d2:", datastore2.Name(), err)
}
