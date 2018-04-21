package routers

import (
	"hackdfw/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/", &controllers.MainController{})
	beego.Router("/v1/user", &controllers.MainController{}, "post:AssignUser")
	beego.Router("/v1/project", &controllers.MainController{}, "get:ProjectList")

}
