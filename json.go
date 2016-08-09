package main

import (
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
	"fmt"
)

type LiveLogModel struct {
	Id         bson.ObjectId `bson:"_id"json:"id"`
	LiveId     string        `bson:"liveid"json:"liveid"`
	LogContent string        `bson:"logcontent"json:"logcontent"`
	Uid        int           `bson:"uid"json:"uid"`
	Ip         string        `bson:"ip"json:"ip"`
	Tx         int           `bson:"tx"json:"tx"`     // 直播上行流量
	Time       int64         `bson:"time"json:"time"` // 日志生成时间,精确到nano,方便日志分析处理
}

func main() {
	buff := `{
	"_id": "579c479d0654340c7a9b9ae4",
		"liveid": "liveid_10016589_zz3VJVUVbKdev",
		"logcontent": "[url:rtmp://pushws.hifun.mobi/live/10016589_zz3VJVUVbKdev?type=1][lIP:192.168.2.133][rIP:111.202.74.130][cT:229(1)][fA:196 fV:255][fAK=0 fVK=0][aCT:0 vCT:0][aF:22 vF:10][aIT:1464289][vIT:1464099][aDT:0 vDT:0][aPT:0 vPT:0][aC:1 vC:0][dA:0 dV:0][bR:290k/s][AVS:190][eSPSPPS:1][aBC:0 vBC:0][t:0][sys am:0, bg:0, pause:0, cameradev:1, micphonedev:1]",
		"uid": 10016589,
		"ip": "111.206.219.101:53308"
}`
	logInfo := LiveLogModel{}
	if err := json.Unmarshal([]byte(buff), &logInfo); err != nil {

	}

	fmt.Printf("livelogModel:%v\n",logInfo)

}
