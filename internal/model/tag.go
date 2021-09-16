package model

import "github.com/jinzhu/gorm"

type Tag struct {
	ID    int    `gorm:"column:id" json:"id" form:"id"`
	Name  string `gorm:"column:name" json:"name" form:"name"`
	State int    `gorm:"column:state" json:"state" form:"state"`
	*Model
}

func (t *Tag) TableName() string {
	return "blog_tag"
}

func (t *Tag) List(db *gorm.DB, page, pageSize int) (*[]Tag, error) {
	var res = []Tag{}
	err := db.Model(t.TableName()).Find(&res).Error
	if err != nil {
		return nil, err
	}
	return &res, nil
}