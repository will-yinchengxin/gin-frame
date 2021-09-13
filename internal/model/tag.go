package model

type Tag struct {
	ID    int64  `gorm:"column:id" json:"id" form:"id"`
	Name  string `gorm:"column:name" json:"name" form:"name"`
	State int64  `gorm:"column:state" json:"state" form:"state"`
}

func (t *Tag) TableName() string {
	return "blog_tag"
}