package main

import (
	_ "baidu-movie/controllers"
	"baidu-movie/models"
	_ "baidu-movie/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	models.RegisterDB()
	beego.BConfig.WebConfig.Session.SessionGCMaxLifetime = 600
}
func main() {
	beego.SetLogger("file", `{"filename":"logs/test.log"}`)
	orm.Debug = true
	//orm.RunSyncdb("default", false, true)

	beego.Run()
}
