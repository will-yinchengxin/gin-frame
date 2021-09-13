package model

type Article struct {
	ID            int64  `gorm:"column:id" json:"id" form:"id"`
	Title         string `gorm:"column:title" json:"title" form:"title"`
	Desc          string `gorm:"column:desc" json:"desc" form:"desc"`
	CoverImageUrl string `gorm:"column:cover_image_url" json:"cover_image_url" form:"cover_image_url"`
	Content       string `gorm:"column:content" json:"content" form:"content"`
	Status        int64  `gorm:"column:status" json:"status" form:"status"`
}

func (a *Article) TableName() string {
	return "blog_article"
}
