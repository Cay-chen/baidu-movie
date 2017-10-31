package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type UpMovieDataControllers struct {
	beego.Controller
}

func (c *UpMovieDataControllers) Post() {
	if c.GetSession("session_val") == nil || c.Ctx.GetCookie("cookie_val") == "" {
		c.Redirect("/", 302)
		c.StopRun()
	}
	id := c.GetString("id")
	name := c.GetString("name")
	subtype := c.GetString("subtype")
	ratings_count := c.GetString("ratings_count")
	genres := c.GetString("genres")
	countries := c.GetString("countries")
	year := c.GetString("year")
	director := c.GetString("director")
	act := c.GetString("act")
	code := c.GetString("code")
	now_num := c.GetString("now_num")
	total_num := c.GetString("total_num")
	other_name := c.GetString("other_name")
	baidu_url := c.GetString("baidu_url")
	img_url := c.GetString("img_url")
	log := c.GetString("log")
	var maps []orm.Params
	serach_sql := "select id from movie where id = " + id + " limit 1;"
	var result map[string]interface{}
	result = make(map[string]interface{})
	o := orm.NewOrm()
	num, err := o.Raw(serach_sql).Values(&maps)
	if err == nil {
		if num == 1 {
			result["msg"] = "该影视已经上传"
		} else {
			sql := "INSERT INTO movie (id,name,subtype,ratings_count,genres,countries,year,director,act,code,now_num,total_num,other_name,baidu_url,img_url,log,up_time) VALUES('" + id + "','" + name + "'," + subtype + ",'" + ratings_count + "','" + genres + "','" + countries + "','" + year + "','" + director + "','" + act + "','" + code + "'," + now_num + "," + total_num + ",'" + other_name + "','" + baidu_url + "','" + img_url + "','" + log + "',now())"
			_, err := o.Raw(sql).Exec()
			if err == nil {
				result["success"] = true
			} else {
				result["msg"] = "插入数据库出现错误"
			}
		}
	} else {
		result["msg"] = "查询该影视数据库出现错误"

	}

	c.Data["json"] = result

	c.ServeJSON()
}
