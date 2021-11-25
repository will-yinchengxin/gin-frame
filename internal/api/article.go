package api

import (
	"context"
	"frame/internal/request"
	"frame/pkg/code"
	"frame/global"
	"github.com/gin-gonic/gin"
	"frame/internal/service"
)

type Article struct {
	ArticleServer  service.Article
}

func (a *Article) List(c *gin.Context)  {
	//
	code.SuccessWithData(c, map[string]string{"name": "will"})
}

func (a *Article) Code(c *gin.Context) {
	code.Error(c, &code.TagLose)
}

func (a *Article) Validator(c *gin.Context) {
	var req request.OnlineList
	if err := global.ReqValidator.ParseJson(c, &req); err != "" {
		code.ValidatorError(c, code.ParamsError.Code, err)
		return
	}

	code.Success(c)
}

func (a *Article) Gorm(c *gin.Context)  {
	var req request.Company
	if err := global.ReqValidator.ParseJson(c, &req); err != "" {
		code.ValidatorError(c, code.ParamsError.Code, err)
		return
	}
	ctx := context.TODO()

	res, codeType := a.ArticleServer.Gorm(req, ctx)
	if codeType.Code != 0 {
		code.Error(c, codeType)
		return
	}
	code.SuccessWithData(c, res)
}