package main

import (
	"flag"
	"fmt"
	"github.com/liukunxin/crontab/master"
	"runtime"
)

var(
	confFile string //配置文件路径
)
//解析命令行参数
func InitArgs()  {
	//master-config ./master.json -xxx 123 -yyy ddd
	//master -h
	flag.StringVar(&confFile,"config","./master.json","指定master.json")
	flag.Parse()
}
//初始化线程
func initEnv()  {
	runtime.GOMAXPROCS(runtime.NumCPU())
}
func main() {

	var(
		err error
	)
	//初始化命令行参数
	InitArgs()
	//初始化线程
	initEnv()

	//加载配置
	if err = master.InitConfig(confFile);err!=nil {
		goto ERR
	}
	//启动API HTTP服务
	if err = master.InitApiServer(); err != nil {
		goto ERR
	}
	//正常退出
	return

ERR:
	fmt.Println(err)
}