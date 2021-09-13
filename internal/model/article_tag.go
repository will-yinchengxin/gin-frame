package model

type ArticleTag struct {
	ID        int64 `gorm:"column:id" json:"id" form:"id"`
	ArticleId int64 `gorm:"column:article_id" json:"article_id" form:"article_id"`
	TagId     int64 `gorm:"column:tag_id" json:"tag_id" form:"tag_id"`
}

func (a *ArticleTag) TableName() string {
	return "blog_article"
}

