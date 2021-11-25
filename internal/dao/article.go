package dao

import (
	"context"
	"frame/global"
	"frame/internal/model"
)

type Article struct {
	ArticleModel model.Article
}

func (a *Article) GetById(id int64, ctx context.Context) (res model.Article, err error) {
	res = a.ArticleModel
	if err = global.DBEngine.Table(a.ArticleModel.TableName()).
		Where("id = ?", id).
		Find(&res).
		Error; err != nil {
		return model.Article{}, err
	}
	return
}
