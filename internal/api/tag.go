package api

import (
	"frame/pkg/app"
	"frame/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Tag struct {
}

func NewTag() Tag {
	return Tag{}
}

func (t *Tag) Get(c *gin.Context)  {

}

func (t *Tag) List(c *gin.Context)  {
	app.NewResponse(c).ToErrorResponse(errcode.ServerError.WithDetails("test err"))
	return
}

func (t *Tag) Create(c *gin.Context)  {

}

func (t *Tag) Update(c *gin.Context)  {

}

func (t *Tag) Delete(c *gin.Context)  {

}