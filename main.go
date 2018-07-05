package main

import (
	//"github.com/guowenshuai/apiproject/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	_ "github.com/guowenshuai/apiproject/routers"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins: false,
		AllowOrigins: []string{
			"http://localhost:8081",
		},
		AllowMethods: []string{
			"PUT", "GET", "DELETE", "POST", "OPTIONS",
		},
		AllowHeaders: []string{
			"Content-Length",
			"Content-Type",
		},
		AllowCredentials: true,
	}))
	beego.Run()
}
