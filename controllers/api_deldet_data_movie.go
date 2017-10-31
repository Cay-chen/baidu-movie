package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type DeleteMovieDataControllers struct {
	beego.Controller
}

func (c *DeleteMovieDataControllers) Post() {
	if c.GetSession("session_val") == nil || c.Ctx.GetCookie("cookie_val") == "" {
		c.Redirect("/", 302)
		c.StopRun()
	}
	id := c.GetString("id")
	delete_sql := "delete from movie where id = '" + id + "'"
	var result map[string]interface{}
	result = make(map[string]interface{})
	o := orm.NewOrm()
	_, err := o.Raw(delete_sql).Exec()
	if err == nil {
		result["success"] = true
	} else {
		result["msg"] = "删除数据库出现错误"
	}

	c.Data["json"] = result
	c.ServeJSON()
}
