package routes

import (
	V1 "frame/internal/api"
	"github.com/gin-gonic/gin"
)

type Tag struct {
	tag V1.Tag
}

func (api *Tag) Initialize(app *gin.Engine) {
	appGroup := app.Group("/v1/api") // .Use(api.MdChannel.CheckSign())
	{
		appGroup.POST("/tags", api.tag.Create)
		appGroup.DELETE("/tags/:id", api.tag.Delete)
		appGroup.PUT("/tags/:id", api.tag.Update)
		appGroup.PATCH("/tags/:id/state", api.tag.Update)
		appGroup.GET("/tags", api.tag.List)
	}
}
