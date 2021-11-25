package routes

import (
	V1 "frame/internal/api"
	"github.com/gin-gonic/gin"
)

type Artists struct {
	// 添加控制器 或 中间件 如: MdChannel.CheckSign()
	article V1.Article
}

func (api *Artists) Initialize(app *gin.Engine) {
	appGroup := app.Group("/v1/api") // .Use(api.MdChannel.CheckSign())
	{
		appGroup.GET("/articles", api.article.List)
	}
}
