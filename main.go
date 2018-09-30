package main

import (
	"test/conf"
	_ "test/routers"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ilibs/gosql"
)

func main() {
	conf.Init()

	//session
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionName = "sessionId"
	beego.BConfig.WebConfig.Session.SessionCookieLifeTime = 0

	//数据库初始化
	configs := make(map[string]*gosql.Config)
	configs["default"] = &conf.Config.DataBase
	gosql.Connect(configs)
	//	注册静态目录
	beego.SetStaticPath("/assets", "assets")
	beego.SetStaticPath("/admin/assets", "assets")
	beego.Run()
}
