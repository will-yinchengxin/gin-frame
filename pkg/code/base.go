package code

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type CodeType struct {
	Code int    `json:"code"`
	Msg  string `json:"message"`
}

type repType struct {
	CodeType
	Data interface{} `json:"data"`
}

func sendResponse(rep CodeType, data interface{}) *repType {
	r := new(repType)
	r.Code = rep.Code
	r.Msg = rep.Msg
	r.Data = data
	return r
}

// 请求成功, 并返回数据
func SuccessWithData(ctx *gin.Context, data interface{}) {
	retData := sendResponse(SUCCESS, data)
	ctx.JSON(http.StatusOK, retData)
	ctx.Abort()
	return
}

// 请求成功
func Success(ctx *gin.Context) {
	retData := sendResponse(SUCCESS, map[string]interface{}{})
	ctx.JSON(http.StatusOK, retData)
	ctx.Abort()
	return
}

// Error 输出错误
func Error(ctx *gin.Context, rep CodeType) {
	retData := sendResponse(rep, map[string]interface{}{})
	ctx.JSON(http.StatusOK, retData)
	ctx.Abort()
	return
}

func ValidatorError(ctx *gin.Context, code int, msg string) {
	retData := sendResponse(CodeType{
		Code: code,
		Msg:  msg,
	}, map[string]interface{}{})
	ctx.JSON(http.StatusOK, retData)
	ctx.Abort()
	return
}
