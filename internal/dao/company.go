package dao

import (
	"context"
	"frame/global"
	"frame/internal/model"
	otgorm "github.com/eddycjy/opentracing-gorm"
	"github.com/jinzhu/gorm"
)

type Company struct {
	CompanyModel model.Company
}

func (c *Company) WithContext(ctx context.Context) *gorm.DB {
	return otgorm.WithContext(ctx, global.DBEngine)
}

func (c *Company) GetById(id int64, ctx context.Context) (res model.Company, err error) {
	// 此处的 debug 与 model.NewDBEngine 中的 debug 设置效果相同
	res = model.Company{}
	if err = c.WithContext(ctx).Table(c.CompanyModel.TableName()).
		Where("id = ?", id).
		Find(&res).
		Error; err != nil {
		return model.Company{}, err
	}
	return
}

func (c *Company) GetList(ctx context.Context) (res []model.Company, err error) {
	if err = c.WithContext(ctx).Table(c.CompanyModel.TableName()).Debug().
		Find(&res).
		Error; err != nil {
		return nil, err
	}
	return
}