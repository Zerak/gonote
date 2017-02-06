package main

import (
	"fmt"
	//"time"
	//"strconv"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"reflect"
	"strconv"
	"time"
)

const (
	MongoDB    = "background"
	Collection = "test"
)

type UserDataMongo struct {
	Id           bson.ObjectId `bson:"_id"json:"id"`
	Uid          int           `bson:"uid"json:"uid"`   // 当前用户id
	Date         int           `bson:"date"json:"date"` // 时间,某天
	Record       []interface{}
	ContinueTime int64 `bson:"continuetime"json:"continuetime"` // 时间,在线时长,秒
}

type Record struct {
	LoginTime  int64
	LogoutTime int64
}

func main() {
	//session, err := mgo.Dial("mongodb://background:lybg0712op@123.56.190.201:27017/background")
	session, err := mgo.Dial("mongodb://192.168.1.25:27017/background")
	if err != nil {
		panic(fmt.Sprintf("mongo dial err:", err))
	}

	fmt.Println("connect success")

	session.Clone()
	defer session.Close()
	collection := session.DB(MongoDB).C(Collection)

	////insert test
	//var datas []interface{}
	//for i := 0; i < 5; i++ {
	//	y, m, d := time.Now().Date()
	//	date, _ := strconv.Atoi(fmt.Sprintf("%04d%02d%02d", y, m, d))
	//	data := &UserDataMongo2{}
	//	data.Id = bson.NewObjectId()
	//	data.Uid = i
	//	data.Date = date
	//
	//	logrec := Record{}
	//	logrec.LoginTime = time.Now().Unix()
	//	logrec.LogoutTime = time.Now().Add(time.Second * 20).Unix()
	//	data.Record = append(data.Record, logrec)
	//	data.ContinueTime += logrec.LogoutTime - logrec.LoginTime
	//
	//	data.Record = append(data.Record, logrec)
	//	data.ContinueTime += logrec.LogoutTime - logrec.LoginTime
	//
	//	datas = append(datas, data)
	//}
	//now := time.Now()
	//err = collection.Insert(datas...)
	//if err != nil {
	//	panic(fmt.Sprintf("insert err :%v\n", err))
	//}
	//fmt.Printf("insert excute time:%v", time.Now().Sub(now))

	// update test
	y, m, d := time.Now().Date() // 当前日期
	date, _ := strconv.Atoi(fmt.Sprintf("%04d%02d%02d", y, m, d))

	mp := make(map[string]string)
	mp["a"] = "a"
	mp["b"] = "a"
	mp["c"] = "a"
	mp["d"] = "a"
	mp["e"] = "a"
	fmt.Printf("mp:f:%v\n", mp["f"])
	for k, v := range mp {
		fmt.Printf("%v:%v\n", k, v)
	}
	fmt.Println()

	results := []UserDataMongo{}
	collection.Find(bson.M{"uid": 2}).All(&results)
	if len(results) > 0 {
		res := results[len(results)-1] // 取最后一条userdata
		recLen := len(res.Record)
		//rec := res.Record[resLen-1].(Record) // 取最后一条Record数据
		rec := res.Record[recLen-1].(bson.M) // 取最后一条Record数据
		fmt.Printf("recType:%v\n", reflect.TypeOf(rec))

		fmt.Println()
		for k, v := range rec {
			fmt.Printf("%v:%v type:k %v, v %v\n", k, v, reflect.TypeOf(k), reflect.TypeOf(v))
		}

		var (
			iLogin  = 0
			iLogout = 0
		)
		login := rec["logintime"].(int64)
		logout := rec["logouttime"].(int64)
		fmt.Printf("login:%v logout:%v logout:%v\n", login, logout, reflect.TypeOf(logout))
		if login != 0 {
			tm := time.Unix(login, 0)
			iLogin, _ = strconv.Atoi(tm.Format("20060102"))
		}
		if logout != 0 {
			tm := time.Unix(logout, 0)
			iLogout, _ = strconv.Atoi(tm.Format("20060102"))
		}
		fmt.Printf("logouttime:%v logintime:%v\n", iLogout, iLogin)

		if iLogout == 0 {
			// 更新logouttime
			con := bson.M{
				"uid": 2, "date": date, "record.logintime": login,
			}
			var updateRec []interface{}
			rec := Record{}
			rec.LoginTime = login
			rec.LogoutTime = time.Now().Unix()
			updateRec = append(updateRec, rec)
			update := bson.M{
				"$set": bson.M{"record.$.logouttime": rec.LogoutTime},
			}
			err = collection.Update(con, update)
			if err != nil {
				panic(fmt.Sprintf("update err :%v\n", err))
			}
			fmt.Println("update ok", rec.LogoutTime)
		}

		fmt.Println()

		fmt.Printf("resuletsLen:%v recLen:%v rec:%v\n", len(results), recLen, rec)
		fmt.Println(res)
	}

	//var updateRec []interface{}
	////for i := 0; i < 3; i++ {
	////	uprec := Record{}
	////	uprec.LoginTime = time.Now().Unix()
	////	uprec.LogoutTime = time.Now().Add(time.Second * 10).Unix()
	////	updateRec = append(updateRec, uprec)
	////
	////	//update := bson.M{
	////	//	"$pushAll": bson.M{"Record": uprec},
	////	//}
	////	//err = collection.Update(con, update)
	////	//if err != nil {
	////	//	panic(fmt.Sprintf("update err :%v\n", err))
	////	//}
	////}
	//update := bson.M{
	//	"$pushAll": bson.M{"Record": updateRec},
	//}
	//err = collection.Update(con, update)
	//if err != nil {
	//	panic(fmt.Sprintf("update err :%v\n", err))
	//}
}
