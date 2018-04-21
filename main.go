package main

import (
	_ "hackdfw/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
	"hackdfw/models"
)

func main() {
	orm.RegisterDriver(`posgtgres`, orm.DRPostgres)
	orm.RegisterDataBase(
		`default`,
		`postgres`,
		`imawnaohofowpo:0bee41441306972ae7ea24310e031c4a15b9806aad23dc36ee575d0ae008fcf2@ec2-54-83-23-91.compute-1.amazonaws.com
/de8rct7mimur2p`)
	orm.DefaultTimeLoc = time.UTC
	orm.RegisterModel(new(models.User))


	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
