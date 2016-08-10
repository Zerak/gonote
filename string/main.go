package main

import (
	"strconv"
	"strings"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"time"
	"live_server/common"
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
	info := LiveLogModel{}
	info.Id = ""
	info.LiveId = "10016614_2S3DB4ZsAKdev"
	info.LogContent = "[url:rtmp://pushws.hifun.mobi/live/10016614_2S3DB4ZsAKdev?type=1][lIP:192.168.2.2][rIP:211.94.114.30][cT:2350(1)][fA:0 fV:0][fAK=0 fVK=0][aCT:0 vCT:0][aF:0 vF:0][aIT:0][vIT:0][aDT:0 vDT:0][aPT:0 vPT:0][aC:0 vC:0][dA:0 dV:0][bR:0123456789B/s][AVS:0][eSPSPPS:0][aBC:0 vBC:0][t:0][sys am:0, bg:0, pause:0, cameradev:1, micphonedev:1]"
	info.Uid = 10016614
	info.Ip = "106.38.167.221:38946"

	WriteLog(info, "addr")
}

func WriteLog(logInfo LiveLogModel, ipaddress string) {
	if logInfo.LiveId != "" && logInfo.LogContent != "" {
		if index := strings.Index(logInfo.LiveId, "?type="); index != -1 {
			logInfo.LiveId = common.SubString(logInfo.LiveId, 0, index)
		}
		logInfo.LiveId = strings.TrimPrefix(logInfo.LiveId, "liveid_")
		logInfo.Ip = ipaddress
		if strings.Contains(logInfo.LogContent, "[bR:") {
			prefix := "bR:"
			brIdx := strings.Index(logInfo.LogContent, "bR:")
			newStr := common.SubString(logInfo.LogContent, brIdx+len(prefix), len(logInfo.LogContent)-brIdx)
			numIdx := 0
			for i, v := range newStr {
				if v >= 48 && v <= 57 { // 0-9 assic 48-57
					continue
				}
				numIdx = i // 数字索引
				break
			}
			spdStr := common.SubString(newStr, 0, numIdx)
			spd := 65535	// 如果是k/s,默认10,即不处理断线,向下兼容

			strTye := common.SubString(newStr, numIdx, len(logInfo.LogContent)-numIdx) // k/s... OR b/s...
			strTye = strings.TrimLeft(strTye, " ")
			if strTye[0] == 'b' {
				spd, _ = strconv.Atoi(spdStr)
			} else if strTye[0] == 'k' {
			}

			logInfo.Tx = spd
			logInfo.Time = time.Now().UnixNano()
			//logModelChan <- logInfo
		} else {
		}

		fmt.Println(logInfo)
	}
}
