package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
)

type GetDoubanDataControllers struct {
	beego.Controller
}

func (c *GetDoubanDataControllers) Get() {
	if c.GetSession("session_val") == nil || c.Ctx.GetCookie("cookie_val") == "" {
		c.Redirect("/", 302)
		c.StopRun()
	}
	url := "https://api.douban.com/v2/movie/subject/" + c.GetString("id")
	req := httplib.Get(url)
	str, err := req.String()
	if err != nil {

	}
	c.Ctx.WriteString(str)
}
