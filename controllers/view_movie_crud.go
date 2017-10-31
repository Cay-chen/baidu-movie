package controllers

import (
	"github.com/astaxie/beego"
)

type ViewMovieCrudControllers struct {
	beego.Controller
}

func (c *ViewMovieCrudControllers) Get() {
	if c.GetSession("session_val") == nil || c.Ctx.GetCookie("cookie_val") == "" {
		c.Redirect("/", 302)
		c.StopRun()
	}
	c.TplName = "moviesCRUD.html"
}
