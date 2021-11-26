package routes

import (
	V1 "frame/internal/api"
	"frame/internal/middlerware"
	"github.com/gin-gonic/gin"
)

type Artists struct {
	// 添加控制器 或 中间件 如: MdChannel.CheckSign()
	article V1.Article

}

func (api *Artists) Initialize(app *gin.Engine) {
	appGroup := app.Group("/v1/api").Use(middlerware.Tracer())
	{
		appGroup.GET("/articles", api.article.List)
		appGroup.GET("/testCode", api.article.Code)
		appGroup.GET("/testValidator", api.article.Validator)
		appGroup.GET("/testGorm", api.article.Gorm)
		appGroup.GET("/testGormAndTracer", api.article.GormAndTracer)
		appGroup.GET("/testRedis", api.article.Redis)
	}
}

