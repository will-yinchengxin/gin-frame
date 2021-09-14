package api

import (
	"frame/global"
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
	// test error
	//app.NewResponse(c).ToErrorResponse(errcode.ServerError.WithDetails("test err"))

	// test validate
	param := struct {
		Name string `json:"name" form:"name" binding:"max=100"`
		State uint8 `form:"state" binding:"oneof=1 2"`
	}{}
	response := app.NewResponse(c)
	valid, errors := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid err cause: %v", errors)
		errRsp := errcode.InvalidParameter.WithDetails(errors.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}
	response.ToResponse(gin.H{})
	return
}

func (t *Tag) Create(c *gin.Context)  {

}

func (t *Tag) Update(c *gin.Context)  {

}

func (t *Tag) Delete(c *gin.Context)  {

}