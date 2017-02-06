package main

import "runtime"
import "log"

func main() {
	test()
}

func test() {
	pc, file, line, ok := runtime.Caller(0)
	log.Println(pc)
	log.Println(file)
	log.Println(line)
	log.Println(ok)
	f := runtime.FuncForPC(pc)
	log.Println(f.Name())

	//pc, file, line, ok = runtime.Caller(0)
	//log.Println(pc)
	//log.Println(file)
	//log.Println(line)
	//log.Println(ok)
	//f = runtime.FuncForPC(pc)
	//log.Println(f.Name())
	//
	//pc, file, line, ok = runtime.Caller(1)
	//log.Println(pc)
	//log.Println(file)
	//log.Println(line)
	//log.Println(ok)
	//f = runtime.FuncForPC(pc)
	//log.Println(f.Name())
}
