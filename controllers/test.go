package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type TestMysqlControllers struct {
	beego.Controller
}

func (c *TestMysqlControllers) Get() {
	o := orm.NewOrm()
	fmt.Println("###########################>", "get")
	//fmt.Println(c.Head())
	var maps1 []orm.Params
	o.Raw("select * from movie where id = ?", "26678594").Values(&maps1)

	fmt.Println(maps1)

	/*res, err := o.Raw("select * from movie where id = ?", "26678594").Exec()
	if err == nil {
		//num, aaa := res.RowsAffected()
		fmt.Println(res)
	}*/

	fmt.Println(c.Ctx.Input.RequestBody)
	fmt.Println(c.GetString("name"))
	c.Ctx.WriteString("get")

}

/*func (c *TestMysqlControllers) Prepare() {
	fmt.Println("###########################>", "Prepare")

}
func (c *TestMysqlControllers) Post() {
	fmt.Println("###########################>", "Post")
	c.Ctx.WriteString("Post")

}
func (c *TestMysqlControllers) Delete() {
	fmt.Println("###########################>", "Delete")
	c.Ctx.WriteString("Delete")

}
func (c *TestMysqlControllers) Put() {
	fmt.Println("###########################>", "Put")
	c.Ctx.WriteString("Put")

}
func (c *TestMysqlControllers) Head() {
	fmt.Println("###########################>", "Head")
	c.Ctx.WriteString("Head")

}
func (c *TestMysqlControllers) Patch() {
	fmt.Println("###########################>", "Patch")
	c.Ctx.WriteString("Patch")

}
func (c *TestMysqlControllers) Options() {
	fmt.Println("###########################>", "Options")
	c.Ctx.WriteString("Options")

}
func (c *TestMysqlControllers) Finish() {
	fmt.Println("###########################>", "Finish")

}
*/
