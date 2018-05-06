package main

import (
	_ "github.com/lanfeust21/AParkhomov/wepapp/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/beego/i18n"
)

func init() {
	orm.RegisterDriver("sqlite3", orm.DRSqlite)
	err := orm.RegisterDataBase("default", "sqlite3", "file:db/sqlite/aparkhomov.db3")
	if err != nil {
		panic(err)
	}
}

func main() {
	beego.AddFuncMap("i18n", i18n.Tr)
	beego.BConfig.EnableGzip = true
	orm.Debug = true
	beego.Run()

}
