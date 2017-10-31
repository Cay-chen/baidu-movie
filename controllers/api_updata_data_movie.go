package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type UpdataMovieDataControllers struct {
	beego.Controller
}

func (c *UpdataMovieDataControllers) Post() {
	if c.GetSession("session_val") == nil || c.Ctx.GetCookie("cookie_val") == "" {
		c.Redirect("/", 302)
		c.StopRun()
	}
	id := c.GetString("id")
	baidu_url := c.GetString("baidu_url")
	now_num := c.GetString("now_num")
	var result map[string]interface{}
	result = make(map[string]interface{})
	update_sql := "update movie set now_num = " + now_num + ",baidu_url ='" + baidu_url + "' where id = '" + id + "'"
	o := orm.NewOrm()
	_, err := o.Raw(update_sql).Exec()
	if err == nil {
		result["success"] = true
	} else {
		result["msg"] = "更新数据库出现错误"
	}
	c.Data["json"] = result

	c.ServeJSON()
}
