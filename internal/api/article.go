package api

import (
	"frame/internal/request"
	"frame/pkg/code"
	"frame/global"
	"github.com/gin-gonic/gin"
)

type Article struct {
}

func (a *Article) List(c *gin.Context)  {
	//
	code.SuccessWithData(c, map[string]string{"name": "will"})
}

func (a *Article) Code(c *gin.Context) {
	code.Error(c, code.TagLose)
}

func (a *Article) Validator(c *gin.Context) {
	var req request.OnlineList
	if err := global.ReqValidator.ParseJson(c, &req); err != "" {
		code.ValidatorError(c, code.ParamsError.Code, err)
		return
	}

	code.Success(c)
}
