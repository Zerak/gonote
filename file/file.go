package main

import (
	//"path/filepath"
	"fmt"
	"log"
	"os"
	"path"
	"time"
)

func main() {
	logPath := "./log"
	logPath = path.Dir(logPath)
	date := time.Now().Format("20060102")
	fpath := logPath + "/" + "livelog/" + date + "/"
	fname := "10019949_1234897asdf" + ".txt"

	e := os.MkdirAll(fpath, 0777)
	if e != nil {
		panic(e)
	}

	//f, err := os.OpenFile(fpath+fname, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	////f, err := os.Create(fpath)
	//if err != nil {
	//	panic(err)
	//}
	//defer f.Close()

	file, err := os.OpenFile(fpath+fname, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		defer file.Close()
		fmt.Printf("open file:%v \nerr:%v\n", fpath, err)
	} else {
		logObj := log.New(file, "\n", log.Ldate|log.Ltime)
		if logObj != nil {
			err := logObj.Output(1, fmt.Sprintf("liveid:%v,uid:%d,ip:%s,chanlen:%d", "liveid_10001", 10019949, "172.100.6.250", 10))
			if err != nil {
				fmt.Printf("logObj.Output err:%v\n", err)
			}
		}
	}

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
