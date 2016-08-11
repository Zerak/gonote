package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type YFRecCallback struct {
	Type        string `json:"type"`         // 录制类型flv,m3u8,hlv,mp4
	StreamEvent string `json:"stream_event"` // 与Type字段内容相同
	OriUrl      string `json:"ori_url"`      // 用户原始流对应的URL
	Domain      string `json:"domain"`       // 对应源流域名
	App         string `json:"app"`          // 挂载点
	Stream      string `json:"stream"`       // 流名
	StartTime   int64  `json:"start_time"`   // 开始录制时间
	StopTime    int64  `json:"stop_time"`    // 停止录制时间
	Duration    int64  `json:"duration"`     // 录制间隔,单位s
	CdnUrl      string `json:"cdn_url"`      // 对应录制文件播放地址
	Size        int    `json:"size"`         // m3u8size值不需要、只有flv的时候才用
}

var (
	param = map[string]string{}
)

func main() {
	decodeStartRecord()
	decodeStopRecord()
}

func decodeStartRecord() {
	jsonStr := `{
	    "type": "flv",
	    "stream_event": "flv",
	    "ori_url": "rtmp://pushyf.hifun.mobi/live/10016593_4VZIi6Cnwxdev",
	    "domain":"send.a.com",
	    "app":"live",
	    "stream":"10016593_4VZIi6Cnwxdev",
	    "start_time": 1470306192,
	    "stop_time": 0,
	    "duration": 0,
	    "cdn_url": "http://"
	}`
	rec := YFRecCallback{}
	if err := json.Unmarshal([]byte(jsonStr), &rec); err != nil {
		fmt.Printf("json unmarshal err:%v\n", err)
	} else {
		param["liveid"] = rec.Stream
		param["url"] = rec.OriUrl
		param["duration"] = strconv.Itoa(int(rec.Duration))
		param["size"] = strconv.Itoa(rec.Size)
		fmt.Printf("decodeStopRecord rec : %v \n%v \n", rec, param)
	}
}

func decodeStopRecord() {
	jsonStr := `{
	    "type": "flv",
	    "stream_event": "flv",
	    "ori_url": "rtmp://pushyf.hifun.mobi/live/10016593_4VZIi6Cnwxdev",
	    "domain":"send.a.com",
	    "app":"live",
	    "stream":"10016593_4VZIi6Cnwxdev",
	    "uri":"hls.a.com/live/hls_bew000/hls_bew000_20160707150625_20160707175817.m3u8",
	    "start_time": 1470306194,
	    "stop_time": 1470306497,
	    "duration": 275,
	    "size":8987799,
	    "cdn_url": "http://hls.a.com/live/hls_bew000/hls_bew000_20160707150625_20160707175817.m3u8"
	}`
	rec := YFRecCallback{}
	if err := json.Unmarshal([]byte(jsonStr), &rec); err != nil {
		fmt.Printf("json unmarshal err:%v\n", err)
	} else {
		param["liveid"] = rec.Stream
		param["url"] = rec.CdnUrl
		param["duration"] = strconv.Itoa(int(rec.Duration))
		param["size"] = strconv.Itoa(rec.Size)
		fmt.Printf("decodeStopRecord rec : %v \n%v \n", rec, param)
	}

}
