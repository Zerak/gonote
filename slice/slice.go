package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type Sql struct {
	id    int
	ip    string
	count int
}

func main() {
	store := make(map[int][]int)
	for i := 0; i < 10; i++ {
		for j := 0; j < 5; j++ {
			store[i] = append(store[i], j)
		}
	}

	store[3] = store[3][1:]

	testSize := make([]int, 0, 1)
	testSize = append(testSize, 1)
	testSize = append(testSize, 2)
	fmt.Printf("size:%v last:%v\n", testSize, testSize[:])

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
		fmt.Printf("key:%v sql:%v type:%v\n", key, sql, reflect.TypeOf(sql))
	}

	// delete of slice
	ts := make([]int, 0)
	fmt.Println("test slice:", ts)
	for i := 0; i < 10; i++ {
		ts = append(ts, i)
	}
	fmt.Println("test slice:", ts)
	uid := 5
	//for k, v := range ts {
	//	if v == uid {
	//		ts = append(ts[:k], ts[k+1:]...)
	//		continue
	//	}
	//	fmt.Println("k:", k, " v:", v)
	//}

	fmt.Println()
	for i := 0; i < len(ts); i++ {
		if ts[i] == uid {
			ts = append(ts[:i], ts[i+1:]...)
			i--
			continue
		}
		fmt.Println("k:", i, " v:", ts[i])
	}
	fmt.Println("test slice:", ts)
}
