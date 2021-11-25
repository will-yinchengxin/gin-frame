package api

import (
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
	//param := struct {
	//	Name string `json:"name" form:"name" binding:"max=100"`
	//	State uint8 `form:"state" binding:"oneof=1 2"`
	//}{}
	//response := app.NewResponse(c)
	//valid, errors := app.BindAndValid(c, &param)
	//if !valid {
	//	global.Logger.Errorf("app.BindAndValid err cause: %v", errors)
	//	errRsp := errcode.InvalidParameter.WithDetails(errors.Errors()...)
	//	response.ToErrorResponse(errRsp)
	//	return
	//}
	//response.ToResponse(gin.H{})
	//return

	// test list-api include service dao

	//var req request.TagListRequest
	//var res = app.NewResponse(c)
	//// 校验参数
	//valid, errors := app.BindAndValid(c, &req)
	//if !valid {
	//	global.Logger.Errorf("GetTagList BindAndValid err cause: %v", errors)
	//	errRsp := errcode.InvalidParameter.WithDetails(errors.Errors()...)
	//	res.ToErrorResponse(errRsp)
	//	return
	//}
	//svc := service.NewService(c.Request.Context())
	////pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	//result, err := svc.GetTagList(&request.TagListRequest{Name: req.Name, State: req.State, Page: req.Page, PageSize: req.PageSize})
	//if err != nil {
	//	global.Logger.Errorf("GetTagList GetTagList err cause: %v", errors)
	//	errRsp := errcode.InvalidParameter.WithDetails(errors.Errors()...)
	//	res.ToErrorResponse(errRsp)
	//	return
	//}
	//res.ToResponseList(result, 2)
}
