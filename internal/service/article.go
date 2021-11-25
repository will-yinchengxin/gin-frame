package service

import (
	"context"
	"frame/internal/dao"
	"frame/internal/request"
	"frame/pkg/code"
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