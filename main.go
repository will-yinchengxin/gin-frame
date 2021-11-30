package main

import (
	"frame/core"
	"frame/global"
	"frame/pkg/mysql"
	"log"
)

func main() {
	// 初始化配置选项
	err := global.SetupSetting()
	if err != nil {
		log.Fatalf("init setting fail, cause %s", err)
	}

	// 初始化日志配置
	err = global.SetupLogger()
	if err != nil {
		log.Fatalf("init Logger fail, cause %s", err)
	}

	// 初始化验证器
	global.SetValidator()

	//初始化数据库连接(这里不再global层做初始化,存在循环依赖的问题)
	global.DBEngine, err = mysql.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		log.Fatalf("init DB fail, cause %s", err)
	}

	// 初始化redis
	err = global.SetRedis()
	if err != nil {
		log.Fatalf("init Redis fail, cause %s", err)
	}

	// 初始化jaegerTrace
	err = global.SetTracer()
	if err != nil {
		log.Fatalf("init JaegerTrace fail, cause %s", err)
	}

	// 初始化 rocketMQ
	err = global.SetRocketMQ()
	if err != nil {
		log.Fatalf("init JaegerTrace fail, cause %s", err)
	}

	//fmt.Println(global.UploadFileSetting.UploadImageAllExts)
	// test log
	//global.Logger.Infof("%s: go-programming-tour-book/%s", "eddycjy")

	print("START SERVER, PORT 8080 \n")
	core.HttpStarter()

	// todo 所有资源连接关闭 及 平滑重启
	//平滑重启或者退出之前执行
	defer func() {
		core.Stop()
	}()
}

