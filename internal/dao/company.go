package dao

import (
	"context"
	"frame/global"
	"frame/internal/model"
)

type Company struct {
	CompanyModel model.Company
}

func (c *Company) GetById(id int64, ctx context.Context) (res model.Company, err error) {
	// 此处的 debug
	res = model.Company{}
	if err = global.DBEngine.Table(c.CompanyModel.TableName()).Debug().
		Where("id = ?", id).
		Find(&res).
		Error; err != nil {
		return model.Company{}, err
	}
	return
}