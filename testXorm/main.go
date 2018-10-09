package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

func main() {
	dn := "mysql"
	sn := "dataSourceName"
	e, err := xorm.NewEngine(dn, sn)
	if err != nil {
		panic(err)
	}
	err = e.Sync()
	if err != nil {
		panic(err)
	}

	sns := []string{"masterDataSourceName", "slave1DataSourceName", "slave2DataSourceName"}
	eg, err := xorm.NewEngineGroup(dn, sns)
	if err != nil {
		panic(err)
	}
	err = eg.Sync()
	if err != nil {
		panic(err)
	}
}
