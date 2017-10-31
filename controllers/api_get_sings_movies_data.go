package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type GetSingsMovieDataController struct {
	beego.Controller
}

func (c *GetSingsMovieDataController) Get() {
	if c.GetSession("session_val") == nil || c.Ctx.GetCookie("cookie_val") == "" {
		c.Redirect("/", 302)
		c.StopRun()
	}
	id := c.GetString("id")
	o := orm.NewOrm()
	sqly := "SELECT * FROM movie WHERE id = '" + id + " '"
	var maps []orm.Params
	num, err := o.Raw(sqly).Values(&maps)
	if err != nil {
		fmt.Println("sssssssssssssssssssss", num, err)

	} else {
		fmt.Println("sssssssssssssssssssss", num, err)

	}
	c.Data["json"] = maps[0]
	c.ServeJSON()

}
func (c *GetSingsMovieDataController) Post() {
	c.Get()
}
