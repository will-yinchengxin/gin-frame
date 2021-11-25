package core

import (
	"frame/global"
	"frame/internal/routes"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func HttpStarter() {
	routeHandler, _ := newRouter()
	s := http.Server{
		Addr:          	global.ServerSetting.HttpPort,
		Handler:        routeHandler,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	err := s.ListenAndServe()
	if err != nil {
		log.Fatalf("start server fail, cause %s", err)
	}
}

// 初始化路由
func newRouter() (*gin.Engine, func()) {
	gin.SetMode(gin.ReleaseMode) // 设置为 release 模式
	engine := gin.Default()
	router := routes.NewRouters
	engine = router().SetupRouter(engine)
	return engine, func() {}
}