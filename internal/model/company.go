package model

type Company struct {
	ID       int64  `gorm:"column:id" json:"id" form:"id"`
	Industry int64  `gorm:"column:industry" json:"industry" form:"industry"`
	Name     string `gorm:"column:name" json:"name" form:"name"`
	Job      string `gorm:"column:job" json:"job" form:"job"`
	UserId   int64  `gorm:"column:user_id" json:"userId" form:"userId"`
}

func (a *Company) TableName() string {
	return "company"
}
