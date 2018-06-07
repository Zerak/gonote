package main

import (
	"encoding/json"
	"fmt"
)

// Response represents request response
type Response struct {
	Result    string `json:"result"`
	TaskID    string `json:"taskid"`
	Desc      string `json:"desc"`
	AuthToken string `json:"auth_token"` // auth sign return
	Status    string `json:"status"`
}

func main() {
	test()
	jsonStr := `
		{
			"result":"ok",
			"expire_time":"1519970099281",
			"auth_token":"0cf902cb60a2149145134ca480fa675bb82de7252be416012aa879ae0254026c"
		}
	`
	res := Response{}
	err := json.Unmarshal([]byte(jsonStr), &res)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}

type QueueMessage struct {
	UID     int    `json:"uid" structs:"uid"`           // 触发行为用户ID
	MsgType int    `json:"msg_type" structs:"msg_type"` // 消息类型ID(必填)
	SubType int    `json:"sub_type" structs:"sub_type"` // 消息类型子ID(必填)
	Rid     int    `json:"rid"`                         // 相关ID
	Content string `json:"content" `                    // 发送内容
}

func test() {
	str := `
	{"msg_type":1,"rid":1508,"sub_type":1,"uid":38}
	`
	s := QueueMessage{}
	err := json.Unmarshal([]byte(str), &s)
	fmt.Println(err)
	fmt.Println(s)
}
