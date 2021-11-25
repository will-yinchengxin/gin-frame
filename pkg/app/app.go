package app

import (
	"frame/pkg/errcode"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CodeType struct {
	Code int    `json:"code"`
	Msg  string `json:"message"`
}

var (
	SUCCESS = CodeType{200, "请求成功"}
	ParamsError    = CodeType{5101, "参数错误"}
	RecordNotFound = CodeType{5102, "记录不存在"}
	ReportError    = CodeType{5103, "数据上报失败"}
	MqError        = CodeType{5104, "mq错误"}
)

type Response struct {
	Ctx *gin.Context
}

func SendResponse(rep CodeType, data interface{}) *RepType {
	r := new(RepType)
	r.Code = rep.Code
	r.Msg = rep.Msg
	r.Data = data
	return r
}

type Pager struct {
	Page      int `json:"page"`
	PageSize  int `json:"page_size"`
	TotalRows int `json:"total"`
}

type RepType struct {
	CodeType
	Data interface{} `json:"data"`
}

func NewResponse(ctx *gin.Context) *Response {
	return &Response{Ctx: ctx}
}

func (r *Response) ToResponse(data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	retData := SendResponse(SUCCESS, data)
	r.Ctx.JSON(http.StatusOK, retData)
	return
}

func (r *Response) ToResponseList(list interface{}, totalRows int) {
	r.Ctx.JSON(http.StatusOK, gin.H{
		"list" : list,
		"pager": Pager{
			Page: GetPage(r.Ctx),
			PageSize: GetPageSize(r.Ctx),
			TotalRows: totalRows,
		},
	})

}

func (r *Response) ToErrorResponse(err *errcode.Error) {
	response := gin.H{"code": err.Code(), "msg": err.Mag()}
	details := err.Details()
	if len(details) > 0 {
		response["details"] = details
	}
	r.Ctx.JSON(err.StatusCode(), response)
}