package main

import (
	"frame/global"
	"frame/internal/model"
	"frame/internal/routes"
	"log"
	"net/http"
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

	//初始化数据库连接(这里不再global层做初始化,存在循环依赖的问题)
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		log.Fatalf("init DB fail, cause %s", err)
	}

	router := routes.NewRouter()

	// test log
	//global.Logger.Infof("%s: go-programming-tour-book/%s", "eddycjy")

	s := http.Server{
		Addr:          	global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

