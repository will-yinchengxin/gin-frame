package api

import (
	"context"
	"frame/consts"
	"frame/global"
	"frame/internal/request"
	"frame/internal/service"
	"frame/pkg/code"
	"github.com/gin-gonic/gin"
	"os"
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

func (a *Article) Gorm(c *gin.Context) {
	var req request.Company
	if err := global.ReqValidator.ParseJson(c, &req); err != "" {
		code.ValidatorError(c, code.ParamsError.Code, err)
		return
	}

	spanFather, _ := c.Get(consts.SpanFather)
	spanFatherCtx := spanFather.(context.Context)
	res, codeType := a.ArticleServer.Gorm(req, spanFatherCtx)
	if codeType.Code != 0 {
		code.Error(c, codeType)
		return
	}
	dir, _ := os.Getwd()
	global.Logger.Info(c, "route Gorm get data success")
	global.Logger.Infof(c, "%s: route Gorm get data success", dir)
	code.SuccessWithData(c, res)
}

func (a *Article) GormAndTracer(c *gin.Context) {
	spanFather, _ := c.Get(consts.SpanFather)
	spanFatherCtx := spanFather.(context.Context)
	res, codeType := a.ArticleServer.GormAndTracer(spanFatherCtx)
	if codeType.Code != 0 {
		code.Error(c, codeType)
		return
	}

	code.SuccessWithData(c, res)
}