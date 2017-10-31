package controllers

import (
	"github.com/astaxie/beego"
)

type ViewUpDataDoubanControllers struct {
	beego.Controller
}

func (c *ViewUpDataDoubanControllers) Get() {
	if c.GetSession("session_val") == nil || c.Ctx.GetCookie("cookie_val") == "" {
		c.Ctx.WriteString("error:请重新登录！")
		c.StopRun()
	}
	c.TplName = "upMovie.html"
}
