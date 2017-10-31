package controllers

import (
	"github.com/astaxie/beego"
	"baidu-movie/utils"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Prepare() {

}

func (c *HomeController) Get() {

	if c.GetSession("session_val") == nil || c.Ctx.GetCookie("cookie_val") == "" {
		c.Redirect("/", 302)
		c.StopRun()
	}
	c.TplName = "home.html"
}
func (c *HomeController) Post() {
	user := c.GetString("user")
	pwd := c.GetString("pwd")
	logins := utils.Md5R32(user+"HXT!56$"+pwd+"8/.dhf") == beego.AppConfig.String("lonin")
	if !logins {
		c.Redirect("/", 302)
		c.StopRun()
	}
	c.SetSession("session_val", utils.Md5R32(user+"0.j8-"))

	c.Ctx.SetCookie("cookie_val", utils.Md5R32(pwd+"6`=k9"), 600, "/")

	c.TplName = "home.html"
}
