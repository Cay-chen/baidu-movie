package controllers

import (
	"github.com/astaxie/beego"
)

type ViewUpDataDoubanControllers struct {
	beego.Controller
}

func (c *ViewUpDataDoubanControllers) Get() {
	if c.GetSession("session_val") == nil || c.Ctx.GetCookie("cookie_val") == "" {
		c.Redirect("/", 302)
		c.StopRun()
	}
	c.TplName = "upMovie.html"
}
