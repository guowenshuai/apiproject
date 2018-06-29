package db

import (
	"github.com/astaxie/beego"
	beeOrm "github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func init() {
	mysqluser, mysqlpwd, mysqlurl, mysqlport, mysqldb :=
		beego.AppConfig.String("mysqluser"),
		beego.AppConfig.String("mysqlpass"),
		beego.AppConfig.String("mysqlurls"),
		beego.AppConfig.String("mysqlport"),
		beego.AppConfig.String("mysqldb")

	dburl := mysqluser + ":" + mysqlpwd + "@tcp(" + mysqlurl + ":" + mysqlport + ")/" + mysqldb

	connIsOk := make(chan bool)
	go registerDB(dburl, connIsOk)
	if isok := <-connIsOk; isok {
		beego.Info("connect db success")
	} else {
		panic("can't connect db")
	}

	// 打开调试模式
	beeOrm.Debug = true

	// 注册驱动
	beeOrm.RegisterDriver("mysql", beeOrm.DRMySQL)
}

func registerDB(dburl string, connOk chan bool) {
	count := 0
	defer func() {
		if err := recover(); err != nil {
			connOk <- false
			close(connOk)
		} else {
			connOk <- true
			close(connOk)
		}
	}()

connDB:
	// 注册数据库
	if count > 60 {
		connOk <- false
		close(connOk)
		panic("connect db timeout")
	}
	beego.Info("try to connect db: ", dburl, " times ", count)
	if err := beeOrm.RegisterDataBase("default", "mysql", dburl+"?charset=utf8, 30, 30"); err != nil {
		beego.Error("register db err: ", err.Error())
		time.Sleep(time.Second * 2)
		count++
		goto connDB
	}

}
