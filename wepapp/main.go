package main

import (
	_ "github.com/lanfeust21/AParkhomov/wepapp/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	orm.RegisterDriver("sqlite3", orm.DRSqlite)
	err:=orm.RegisterDataBase("default", "sqlite3", "file:db/sqlite/aparkhomov.db3")
	if err!=nil{
		panic(err)
	}
}

func main() {
	beego.Run()
}

