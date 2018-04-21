package controllers

import (
	"github.com/astaxie/beego"

)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *MainController) ProjectList() {

	var region string
	c.Ctx.Input.Bind(&region, "region")
	lp := data_model.GetMarketList(region)
	c.Data["json"] = lp
	c.ServeJSON()

}