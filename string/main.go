package main

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"live_server/common"
	"strconv"
	"strings"
	"time"
)

type LiveLogModel struct {
	Id         bson.ObjectId `bson:"_id"json:"id"`
	LiveId     string        `bson:"liveid"json:"liveid"`
	LogContent string        `bson:"logcontent"json:"logcontent"`
	Uid        int           `bson:"uid"json:"uid"`
	Ip         string        `bson:"ip"json:"ip"`
	Tx         int64         `bson:"tx"json:"tx"`     // 直播上行流量
	Time       int64         `bson:"time"json:"time"` // 日志生成时间,精确到nano,方便日志分析处理
}

func main() {
	info := LiveLogModel{}
	info.Id = ""
	info.LiveId = "10016614_2S3DB4ZsAKdev"
	//info.LogContent = "[url:rtmp://pushws.hifun.mobi/live/10016614_2S3DB4ZsAKdev?type=1][lIP:192.168.2.2][rIP:211.94.114.30][cT:2350(1)][fA:0 fV:0][fAK=0 fVK=0][aCT:0 vCT:0][aF:0 vF:0][aIT:0][vIT:0][aDT:0 vDT:0][aPT:0 vPT:0][aC:0 vC:0][dA:0 dV:0][bR:0.012345678901234567890. k/s][AVS:0][eSPSPPS:0][aBC:0 vBC:0][t:0][sys am:0, bg:0, pause:0, cameradev:1, micphonedev:1]"
	info.LogContent = "[dA:0 dV:0][bR:0.123k/s][AVS:0][eSPSPPS:0][aBC:0 vBC:0][t:0][sys am:0, bg:0, pause:0, cameradev:1, micphonedev:1]"
	info.Uid = 10016614
	info.Ip = "106.38.167.221:38946"

	now := time.Now().UnixNano()
	WriteLog(info, "addr")
	fmt.Printf("process time:%v us\n",(time.Now().UnixNano() - now) / 1000)

	//now = time.Now().UnixNano()
	//WriteLog2(info, "addr")
	//fmt.Printf("process time:%v us\n",(time.Now().UnixNano() - now) / 1000)
}

func WriteLog(logInfo LiveLogModel, ipaddress string) {
	fmt.Printf("writeLog...\n")
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
				if v == 46 {	// . assic 46
					continue
				}
				numIdx = i // 数字索引
				break
			}
			spdStr := common.SubString(newStr, 0, numIdx)
			var spd int64

			strTye := common.SubString(newStr, numIdx, len(logInfo.LogContent)-numIdx) // k/s... OR b/s...
			strTye = strings.TrimLeft(strTye, " ")
			if strTye[0] == 'b' {
				ss, _ := strconv.Atoi(spdStr)
				spd = (int64)(ss)
			} else if strTye[0] == 'k' {
				ss, _ := strconv.ParseFloat(spdStr, 64)
				spd = (int64)(ss * 1000)
			}

			fmt.Printf("type[%v] spd:%v\n", string(strTye[0]), spd)

			logInfo.Tx = spd
			logInfo.Time = time.Now().UnixNano()
			//logModelChan <- logInfo
		} else {
		}

		fmt.Println(logInfo)
	}
}

func WriteLog2(logInfo LiveLogModel, ipaddress string) {
	fmt.Printf("\nwriteLog2...\n")
	if logInfo.LiveId != "" && logInfo.LogContent != "" {
		if index := strings.Index(logInfo.LiveId, "?type="); index != -1 {
			logInfo.LiveId = common.SubString(logInfo.LiveId, 0, index)
		}
		logInfo.LiveId = strings.TrimPrefix(logInfo.LiveId, "liveid_")
		logInfo.Ip = ipaddress
		if strings.Contains(logInfo.LogContent, "[bR:") && strings.Contains(logInfo.LogContent, " k/s") {
			prefix := "bR:"
			brIdx := strings.Index(logInfo.LogContent, "bR:")
			ksidx := strings.Index(logInfo.LogContent, " k/s")
			str := common.SubString(logInfo.LogContent, brIdx+len(prefix), ksidx-brIdx-len(prefix))
			spd, _ := strconv.ParseFloat(str, 64)
			fmt.Printf("spd:%v\n", spd)

			logInfo.Tx = (int64)(spd * 1000)
			logInfo.Time = time.Now().UnixNano()
			//logModelChan <- logInfo
		} else {
		}

		fmt.Println(logInfo)
	}
}
