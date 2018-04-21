package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
)

type UsersController struct {
	beego.Controller
}

/*
FN
LN
email
pass

 */

func (r *UsersController) Post()  {
	fmt.Print(r["first_name"])
}
