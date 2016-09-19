package main

import (
	"fmt"
	"encoding/json"
	"strconv"
)

func main() {
	reData := make(map[string]map[string]string)
	uData := make(map[string]string)
	//uData["nickname"] = "zhangshan"
	//uData["avarta"] = "avaaddr"
	//reData["100"] = uData

	for i := 0; i < 3; i++ {
		uData["nickname"] = "zhangsha" + strconv.Itoa(i)
		uData["avata"] = "advaddr" + strconv.Itoa(i)
		reData[strconv.Itoa(i)] = uData
	}

	fmt.Println(reData)
	by , err := json.Marshal(reData)
	if err != nil {
		fmt.Printf("err:%v\n",err)
	}
	fmt.Printf("data str:%v\n",string(by))
	fmt.Printf("data bye:%v\n",by)
}
