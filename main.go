package main

import (
	_ "hackdfw/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
	_ "github.com/lib/pq"

	"os"

)

func main() {
	orm.RegisterDriver(`postgres`, orm.DRPostgres)
	orm.RegisterDataBase(
		`default`,
		`postgres`,
		`postgres://imawnaohofowpo:0bee41441306972ae7ea24310e031c4a15b9806aad23dc36ee575d0ae008fcf2@ec2-54-83-23-91.compute-1.amazonaws.com/de8rct7mimur2p`)
	orm.DefaultTimeLoc = time.UTC

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	//beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
	//	AllowOrigins:     []string{"*"},
	//	AllowMethods:     []string{"PUT", "PATCH", "GET", "DELETE"},
	//	AllowHeaders:     []string{"Origin"},
	//	ExposeHeaders:    []string{"Content-Length"},
	//	AllowCredentials: false,
	//}))

	beego.Run(":" + port)
}
