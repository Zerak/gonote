package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"yb_infra/code.google.com/p/mahonia"
)

func main() {
	flag.Parse()

	arg := flag.String("file", "", "the file name to be parse")
	if *arg == "" {
		fmt.Println("err arg")
		return
	}
	fmt.Println(*arg)
	return

	result, err := os.OpenFile("result.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("open result err:", err)
		return
	}
	defer result.Close()

	sf, err := os.OpenFile("search.txt", os.O_RDONLY, 0666)
	if err != nil {
		return
	}
	defer sf.Close()

	count := 0
	buf := bufio.NewReader(sf)
	for {
		line, _, err := buf.ReadLine()
		if err != nil {
			fmt.Println("read line err:", err)
			break
		}
		str := string(line)
		idxKey := strings.Index(str, "post[map[keyword:[")
		if idxKey != -1 {
			keyword := str[idxKey+18 : idxKey+18+12]
			keyword = strings.TrimRight(keyword, "]")
			idxScore := strings.Index(keyword, "]")
			if idxScore != -1 {
				fmt.Println(keyword)
				keyword = keyword[0:idxScore]
				fmt.Println(keyword)
			}

			dec := mahonia.NewDecoder("utf-8")
			kw := dec.ConvertString(keyword)
			n, err := result.WriteString(kw)
			if err != nil {
				fmt.Println("write err:", n, err)
			}
			result.WriteString("\n")
			//err := logObj.Output(2, keyword)
			//if err != nil {
			//	fmt.Println("out err:", err)
			//}
		} else {
			//fmt.Println(idxKey, " str:", str)
			count++
		}
	}
	fmt.Println("-1 count:", count)
	return

	//logPath := "./log"
	//logPath = path.Dir(logPath)
	//date := time.Now().Format("20060102")
	//fpath := logPath + "/" + "livelog/" + date + "/"
	//fname := "10019949_1234897asdf" + ".txt"
	//
	//e := os.MkdirAll(fpath, 0777)
	//if e != nil {
	//	panic(e)
	//}
	//
	////f, err := os.OpenFile(fpath+fname, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	//////f, err := os.Create(fpath)
	////if err != nil {
	////	panic(err)
	////}
	////defer f.Close()
	//
	//file, err := os.OpenFile(fpath+fname, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	//if err != nil {
	//	defer file.Close()
	//	fmt.Printf("open file:%v \nerr:%v\n", fpath, err)
	//} else {
	//	logObj := log.New(file, "\n", log.Ldate|log.Ltime)
	//	if logObj != nil {
	//err := logObj.Output(1, fmt.Sprintf("liveid:%v,uid:%d,ip:%s,chanlen:%d", "liveid_10001", 10019949, "172.100.6.250", 10))
	//		if err != nil {
	//			fmt.Printf("logObj.Output err:%v\n", err)
	//		}
	//	}
	//}

	//dir, file := filepath.Split("./file.go")
	//if dir != "" && dir != "." {
	//	if err := os.MkdirAll(dir,os.FileMode(0775)); err != nil{
	//		fmt.Println("mkdir err")
	//		return
	//	}
	//	fmt.Printf("dir:%v file:%v\n", dir, file)
	//} else {
	//	fmt.Println("dir empty")
	//}
}
