package main

import (
	_ "pmt/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	mysqluser := beego.AppConfig.String("mysqluser")
	mysqlpass := beego.AppConfig.String("mysqlpass")
	mysqlhost := beego.AppConfig.String("mysqlurls")
	mysqlport := beego.AppConfig.String("mysqlport")
	mysqldb   := beego.AppConfig.String("mysqldb")
	mysql_info := mysqluser+":"+mysqlpass+"@tcp("+mysqlhost+":"+mysqlport+")/"+mysqldb

	orm.RegisterDataBase("default", "mysql", mysql_info)
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/static"] = "static"
	}
	beego.Run()
}

