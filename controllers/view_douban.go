package controllers

import (
	"github.com/astaxie/beego"
)

type ViewDoubanControllers struct {
	beego.Controller
}

func (c *ViewDoubanControllers) Get() {
	c.TplName = "hehe.html"
}
