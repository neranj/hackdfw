package routers

import (
	"hackdfw/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router(`users`, &controllers.UsersController{})
}
