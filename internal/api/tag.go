package api

import (
	"github.com/gin-gonic/gin"
)

type Tag struct {
}

func (t *Tag) List(c *gin.Context)  {
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
