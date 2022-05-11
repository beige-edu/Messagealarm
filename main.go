package main

import (
	_ "PrometheusAlert/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

func main() {
	orm.Debug = true
	logtype := beego.AppConfig.String("logtype")
	if logtype == "console" {
		_ = logs.SetLogger(logtype)
	} else if logtype == "file" {
		_ = logs.SetLogger(logtype, `{"filename":"`+beego.AppConfig.String("logpath")+`"}`)
	}
	logs.Info("[main] 当前版本（Version）4.6.1")
	//model.MetricsInit()
	//beego.Handler("/metrics", promhttp.Handler())
	beego.Run()
}
