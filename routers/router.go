package routers

import (
	"baidu-movie/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.LoginControllers{})
	beego.Router("/home", &controllers.HomeController{})
	beego.Router("/api/getdata", &controllers.GetMovieDataController{})
	beego.Router("/view/datacrud", &controllers.ViewMovieCrudControllers{})
	beego.Router("/test/mysql", &controllers.TestMysqlControllers{})
	beego.Router("/api/getdouban", &controllers.GetDoubanDataControllers{})
	beego.Router("/test/douban", &controllers.ViewDoubanControllers{})
	beego.Router("/api/upmovie", &controllers.UpMovieDataControllers{})
	beego.Router("/view/updatamovie", &controllers.ViewUpDataDoubanControllers{})
	beego.Router("/api/singsdata", &controllers.GetSingsMovieDataController{})
	beego.Router("/api/deletedata", &controllers.DeleteMovieDataControllers{})
	beego.Router("/api/update", &controllers.UpdataMovieDataControllers{})

}
