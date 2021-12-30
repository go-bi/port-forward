package main

import (
	"flag"
	"fmt"
	logs "github.com/tavenli/gologs"
)

func main(){

	flag.Parse()

	fmt.Println("----------------------")

	//日志级别："emergency", "alert", "critical", "error", "warning", "notice", "info", "debug"
	logs.SetLogger(logs.AdapterConsole, `{"level":7}`)
	//logs.SetLogger(logs.AdapterFile, `{"filename":"app.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10}`)

	//logs.SetLogger(logs.AdapterMultiFile, `{"filename":"app.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10,"separate":["error", "warning", "info"]}`)
	logs.SetLogger(logs.AdapterMultiFile, `{"filename":"app.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10,"separate":["error"]}`)

	//输出文件名和行号
	logs.EnableFuncCallDepth(true)
	logs.SetLogFuncCallDepth(3)
	//为了让日志输出不影响性能，开启异步日志
	logs.Async()

	logs.Debug("----------------------")
	logs.Info("iiiiiiiiiiiiii")
	logs.Warn("wwwwwwwww")
	logs.Error("eeeeeeeeeeeeeeee")

	for {

	}
}
