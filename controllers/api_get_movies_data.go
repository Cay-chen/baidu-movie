package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/cay/utils"
	"strconv"
	"strings"
)

type GetMovieDataController struct {
	beego.Controller
}

func (c *GetMovieDataController) Get() {
	if c.GetSession("session_val") == nil || c.Ctx.GetCookie("cookie_val") == "" {
		c.Redirect("/", 302)
		c.StopRun()
	}
	num1, _ := strconv.Atoi(c.GetString("page"))
	num2, _ := strconv.Atoi(c.GetString("rows"))
	page := utils.If3(c.GetString("page") != "", num1, 1).(int)
	rows := utils.If3(c.GetString("rows") != "", num2, 10).(int)
	offset := (page - 1) * rows
	o := orm.NewOrm()
	var maps1 []orm.Params
	var maps []orm.Params
	mtype := c.GetString("mtype")
	mval := strings.Trim(c.GetString("mval"), " ")
	if mtype == "" {
		_, err := o.Raw("select count(*) as count from movie").Values(&maps1)
		var total interface{}
		if err != nil {
			fmt.Println("数据库连接错误------------", err)
			_, err := o.Raw("select count(*) as count from movie").Values(&maps1)
			fmt.Println("再次连接数据库后------------", err)
			total = maps1[0]["count"]
		} else {
			total = maps1[0]["count"]
		}

		sqly := "SELECT * FROM movie limit " + strconv.Itoa(offset) + "," + strconv.Itoa(rows)
		o.Raw(sqly).Values(&maps)
		var backResult map[string]interface{}
		backResult = make(map[string]interface{})
		backResult["total"] = total.(string)
		backResult["rows"] = maps
		c.Data["json"] = backResult
		c.ServeJSON()
	} else {
		sql_where := mtype + " like '%" + mval + "%'"
		sql_count := "select count(*) as count from movie where " + sql_where
		_, err := o.Raw(sql_count).Values(&maps1)
		var total interface{}
		if err != nil {
			total = page
		} else {
			total = maps1[0]["count"]
		}
		sqly := "SELECT * FROM movie where " + sql_where + " limit " + strconv.Itoa(offset) + "," + strconv.Itoa(rows)
		o.Raw(sqly).Values(&maps)
		var backResult map[string]interface{}
		backResult = make(map[string]interface{})
		backResult["total"] = total.(string)
		backResult["rows"] = maps
		c.Data["json"] = backResult
		c.ServeJSON()

	}

}
func (c *GetMovieDataController) Post() {
	if c.GetSession("session_val") == nil || c.Ctx.GetCookie("cookie_val") == "" {
		c.Redirect("/", 302)
		c.StopRun()
	}
	c.Get()
}
