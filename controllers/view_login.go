package controllers

import (
	"github.com/astaxie/beego"
)

type LoginControllers struct {
	beego.Controller
}

func (c *LoginControllers) Get() {
	isExit := c.GetString("exit")
	if isExit == "true" {
		c.DelSession("session_val")
		c.Ctx.SetCookie("cookie_val", "", -1, "/")
	}
	c.TplName = "login.html"
}
