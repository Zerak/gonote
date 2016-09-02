package main

import (
	"reflect"
	"strings"
	"strconv"

	"lvstatistics/statistics/model"
	"lvstatistics/common"
	"fmt"
)

func LoadTableStruct(st interface{}) (fields map[string]model.TagStruct) {
	s := reflect.TypeOf(st).Elem()
	fields = make(map[string]model.TagStruct, 0)
	for i := 0; i < s.NumField(); i++ {
		tag := common.ReplaceSpace(string(s.Field(i).Tag))
		if tag != "" {
			fl := strings.Split(tag, " ")
			field := model.TagStruct{}
			for _, l := range fl {
				sl := strings.Split(l, "=")
				if sl[0] == "param" {
					field.Param = sl[1]
				}

				if sl[0] == "field" {
					field.Field = sl[1]
				}

				if sl[0] == "dbtype" {
					field.Dbtype = sl[1]
				}

				if sl[0] == "dblen" {
					field.Dblen, _ = strconv.Atoi(sl[1])
				}
			}

			fields[field.Param] = field
		}
	}
	return fields
}

func main() {
	//var baseTableFields map[string]model.TagStruct
	baseTableFields := LoadTableStruct(new(model.LoginBaseTable))
	for k,v := range baseTableFields {
		fmt.Printf("key:%v v:%v vType:%v\n",k,v,reflect.TypeOf(v))
	}
}