package api

import (
	"frame/pkg/app"
	"github.com/gin-gonic/gin"
)

type Article struct {
}

func (a *Article) List(c *gin.Context)  {
	var res = app.NewResponse(c)
	res.ToResponse(map[string]string{"name":"will"})
}