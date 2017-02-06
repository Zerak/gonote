package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"live_common_lib/common"
)

func main() {
	Test()
	return
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
	by, err := json.Marshal(reData)
	if err != nil {
		fmt.Printf("err:%v\n", err)
	}
	fmt.Printf("data str:%v\n", string(by))
	fmt.Printf("data bye:%v\n", by)
}

func Test() {
	// 网速预取post
	type pubParam struct {
		UserName     string   `json:"user_name"`
		CheckCode    string   `json:"check_code"`
		FetchOption  string   `json:"fetchOption"`
		NeedFeedback string   `json:"need_feedback"`
		UrlList      []string `json:"url_list"`
	}
	pub := pubParam{}
	pub.UserName = "51renzhen"
	pw := "HIfun866"
	str := fmt.Sprintf("%s%s%s%s", time.Now().Format("20060102"), pub.UserName, "chinanetcenter", pw)
	fmt.Printf("post str:%v\n", str)
	code := common.Md5(str)
	pub.CheckCode = code
	pub.FetchOption = "Y"
	pub.NeedFeedback = "0"
	pub.UrlList = append(pub.UrlList, "http://10019949_12321349876asdfjkl.flv")
	pub.UrlList = append(pub.UrlList, "http://10019949_12345789asfd.m3u8")
	bye, err := json.Marshal(&pub)
	if err != nil {
		fmt.Printf("wangsu post marshal jason err:%v\n", err)
	} else {
		fmt.Printf("post data:%v\n", string(bye))
		pb, err := common.HttpPost(bye, "http://cm.chinanetcenter.com/CM/cm-publish!json.do")
		if err != nil {
			fmt.Printf("post err:%v\n", err)
		}
		fmt.Printf("post return:%v\n", string(pb))
	}
}
