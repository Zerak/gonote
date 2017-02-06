package main

import (
	"encoding/json"
	"fmt"
)

type JhFieInfo struct {
	Key    string  `json:"key"`
	Url    string  `json:"url"`
	Rate   float64 `json:"rate"`
	Label  int     `json:"label"`
	Review bool    `json:"review"`
	Err    string  `json:"error"`
}
type PornDetect struct {
	Code     string      `json:"code"`
	Msg      string      `json:"detectId"`
	FileList []JhFieInfo `json:"fileList"`
}
type LiveWangSuJh struct {
	DetectId   string `json:"detectId"`
	Cmd        string `json:"cmd"`
	StreamName string `json:"streamname"`
	Bucket     string `json:"bucket"`
	IsZip      int    `json:"isZip"`
	PornDetect `json:"pornDetect"`
}

func main() {
	jsonStr := `{
	    "detectId": "<detectId>",
	    "cmd": "<cmd>",
	    "streamname": "<streamname>",
	    "bucket": "<bucket>",
	    "isZip": 1,
	    "pornDetect": {
		"code": "<ResultCodeint>",
		"message": "<ResultMessagestring>",
		"fileList": [
		    {
			"key": "<keystring>",
			"url": "<urlstring>",
			"rate": 0.2,
			"label": 1,
			"review": false,
			"error": "error string"
		    }
		]
	    }
	}`

	data := LiveWangSuJh{}
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		fmt.Printf("json unmarshal err:%v", err)
	} else {
		fmt.Println(data)
	}
}
