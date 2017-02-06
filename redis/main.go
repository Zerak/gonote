package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"live_common_lib/qqredis"
)

var RedisClient qqredis.RedisCache

//主程序
func main() {
	fmt.Print(".")
	fmt.Print(".")
	RedisClient = qqredis.NewRedisCache("19f41974e7d74455.m.cnbja.kvstore.aliyuncs.com:6379",
		"19f41974e7d74455:mvnaoNnaflo229jfa", 2000)
	fmt.Printf("connect ok\n")

	//var filename = "/tmp/uids_realname.txt"
	//var filename = "/tmp/uid_realname_all.txt"
	//var filename = "/tmp/check.txt"
	var filename = "/tmp/uid_realname_all_1.txt"
	//filename = "/Users/chengb/workarea/server/src_new/live_server/trunk/src/live_server/web/conf/uids.txt"
	f1, err := os.OpenFile(filename, os.O_RDONLY, 0666)
	if err != nil {
		panic(err)
	}

	inputReader := bufio.NewReader(f1)
	uidArray := [50000]string{}
	index := 0
	for {
		inputString, readerError := inputReader.ReadString('\n')
		if readerError == nil {
			inputString = strings.Replace(inputString, "\n", "", 1)
			uidArray[index] = inputString
		} else {
			if readerError == io.EOF {
				break
			}
		}

		//fmt.Println(inputString)
		index = index + 1
	}

	f1.Close()

	count := 0
	failCount := 0
	for index := 0; index < len(uidArray); index += 1 {
		uid := uidArray[index]
		if len(uid) == 0 {
			continue
		}

		uidKey := "live_u:" + uid
		//_, err := RedisClient.Hsetstring(uidKey, "IsRealname", "1")
		//if err != nil {
		//	fmt.Printf("update user:%v IsRealname err:%v\n", uid, err)
		//	continue
		//}
		//
		//re, err := RedisClient.Hget(uidKey, "IsRealname")
		//if err != nil || re != "1" {
		//	failCount++
		//	fmt.Println(fmt.Sprintf("get uid:%v realname err:%v re:%v", uid, err, re))
		//	continue
		//}
		//count++
		//fmt.Print(".")

		str, _ := RedisClient.Hget(uidKey, "IsRealname")
		val, err := strconv.Atoi(str)
		if err != nil {
			failCount++
			fmt.Println(fmt.Sprintf("uid:%v convert val:%v err:%v", uid, val, err))
			continue
		}
		if val == 0 || val > 3 {
			count++
			fmt.Println(fmt.Sprintf("uid:%v realname:%v", uid, str))
			fmt.Println(fmt.Sprintf("%v", uid))
		}

		time.Sleep(time.Microsecond * 100)
	}
	fmt.Println(fmt.Sprintf("\ntotal:%v fail:%v", count, failCount))
}
