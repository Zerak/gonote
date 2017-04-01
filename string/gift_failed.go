package main

import (
	"fmt"
	"strconv"
	//"io"
	"bufio"
	"io"
	"live_common_lib/qqredis"
	"os"
	"strings"
)

var RedisClient qqredis.RedisCache

//主程序
func main() {
	RedisClient = qqredis.NewRedisCache("19f41974e7d74455.m.cnbja.kvstore.aliyuncs.com:6379",
		"19f41974e7d74455:mvnaoNnaflo229jfa", 2000)
	fmt.Printf("connect ok\n")

	var filename = "/tmp/uids_gift.txt"
	//filename = "/Users/chengb/workarea/server/src_new/live_server/trunk/src/live_server/web/conf/uids.txt"
	f1, _ := os.OpenFile(filename, os.O_RDONLY, 0666)

	inputReader := bufio.NewReader(f1)
	uidArray := map[string]string{}
	index := 0
	for {
		inputString, readerError := inputReader.ReadString('\n')
		if readerError == nil {
			inputString = strings.Replace(inputString, "\n", "", 1)
			str := strings.Split(inputString, ",")
			uidArray[str[0]] = str[1]

		} else {
			if readerError == io.EOF {
				break
			}
		}

		fmt.Println(inputString)
		index = index + 1
	}

	f1.Close()

	for k, v := range uidArray {
		fmt.Printf("k=%v, v=%v\n", k, v)
		uid := k
		value, _ := strconv.Atoi(v)

		if len(uid) == 0 {
			continue
		}

		key := fmt.Sprintf("live_u:%s", uid)
		{
			amount, err := RedisClient.Hget(key, "ConsumptionAmount")
			if err != nil {
				fmt.Printf("get user:%v ConsumptionAmount err:%v\n", uid, err)
				continue
			}
			fmt.Printf("amount:%v\n", amount)

			nextAmount, _ := strconv.Atoi(amount)
			nextAmount = nextAmount + value
			retStr, err := RedisClient.Hsetstring(key, "ConsumptionAmount", strconv.Itoa(nextAmount))
			if err != nil {
				fmt.Printf("update user:%v WithdrawalsAmount err:%v\n", uid, err)
				continue
			} else {
				str, _ := RedisClient.Hgetall(key)
				for _, v := range str {
					fmt.Println("redis hset = " + retStr + " hget = " + v)
				}
			}
		}
	}
}
