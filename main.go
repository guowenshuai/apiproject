package main

import (
	//"github.com/guowenshuai/apiproject/routers"
	"github.com/astaxie/beego"
	_ "github.com/guowenshuai/apiproject/routers"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
