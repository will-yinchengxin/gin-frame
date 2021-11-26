package service

import (
	"context"
	"frame/internal/dao"
	"frame/internal/request"
	"frame/pkg/code"
	"github.com/jinzhu/copier"
)

type Article struct {
	ArticleDao dao.Article
	CompanyDao dao.Company
}

func (a *Article) Gorm(req request.Company, ctx context.Context) (finalData interface{}, codeType *code.CodeType) {
	res, err := a.CompanyDao.GetById(req.ID, ctx)
	if err != nil {
		return finalData, code.MqError
	}
	return res, &code.CodeType{}
}

func (a *Article) GormAndTracer(req request.PageInfo, ctx context.Context) (finalData map[string]interface{}, codeType *code.CodeType) {
	reqPage := request.PageInfo{}
	copier.Copy(&reqPage, &req)

	res, count, err := a.CompanyDao.GetList(reqPage, ctx)
	if err != nil {
		return finalData, code.MqError
	}
	finalData = GetFinalData(req, count, res)
	return finalData, &code.CodeType{}
}