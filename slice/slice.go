package main

import (
	"fmt"
	"strconv"
	"reflect"
)

type Sql struct {
	id    int
	ip    string
	count int
}

func main() {
	var cacheSql []Sql
	for i := 0; i < 10; i++ {
		sql := Sql{}
		sql.id = i
		sql.ip = "sql" + strconv.Itoa(i)
		sql.count = i
		cacheSql = append(cacheSql, sql)
	}

	fmt.Println("befor cacheSql :", cacheSql)

	cacheSql = append(cacheSql[0:0])
	fmt.Println("after cacheSql :", cacheSql)

	for i := 0; i < 10; i++ {
		sql := Sql{}
		sql.id = i
		sql.ip = "sql" + strconv.Itoa(i)
		sql.count = i
		cacheSql = append(cacheSql, sql)
	}
	fmt.Println("reinsert cacheSql :", cacheSql)

	for key, sql := range cacheSql {
		fmt.Printf("sql:%v sqlId:%v ip:%v count:%v\n", sql, sql.id, sql.ip, sql.count)
		fmt.Printf("key:%v sql:%v type:%v\n",key, sql, reflect.TypeOf(sql))
	}
}
