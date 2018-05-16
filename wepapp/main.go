package main

import (
	_ "github.com/lanfeust21/AParkhomov/wepapp/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/beego/i18n"
	"html/template"
	"fmt"
	"regexp"
)

var re = regexp.MustCompile("\\?")

func init() {
	orm.RegisterDriver("sqlite3", orm.DRSqlite)
	err := orm.RegisterDataBase("default", "sqlite3", "file:db/sqlite/aparkhomov.db3")
	if err != nil {
		panic(err)
	}
}

func Link(link interface{}, textAnchor interface{}, querystring string) template.HTML {
	source := ""
	switch link.(type) {
	case int:
		source = fmt.Sprintf("%d", link.(int))
	case float64:
		source = fmt.Sprintf("%2.02", link.(float64))
	case string:
		source = link.(string)
	}

	text := ""
	switch textAnchor.(type) {
	case int:
		text = fmt.Sprintf("%d", textAnchor.(int))
	case float64:
		text = fmt.Sprintf("%2.02", textAnchor.(float64))
	case string:
		text = textAnchor.(string)
	}
	ret := ""
	// if source contain ?
	if re.MatchString(source) {
		ret = fmt.Sprintf("<a href=\"%s&%s\">%s</a>", source, querystring, text)
	} else {
		ret = fmt.Sprintf("<a href=\"%s?%s\">%s</a>", source, querystring, text)
	}
	return template.HTML(ret)
}

func main() {
	beego.AddFuncMap("i18n", i18n.Tr)
	beego.BConfig.EnableGzip = true
	orm.Debug = true

	beego.AddFuncMap("Link", Link)

	beego.Run()

}
