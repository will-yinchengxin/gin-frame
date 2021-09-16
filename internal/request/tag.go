package request

type CountTagRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State int    `form:"state,default=1" binding:"oneof=0 1"`
}

type TagListRequest struct {
	Name     string `form:"name" binding:"max=100"`
	State    int    `form:"state,default=1" binding:"oneof=0 1"`
	Page     int    `form:"page,default=1" binding:"required"`
	PageSize int    `form:"pageSize,default=10" binding:"required"`
}

type CreateTagRequest struct {
	Name     string `form:"name" binding:"max=100"`
	CreateBy string `form:"create_by" binding:"required,min=3,max=100"`
	State    int    `form:"state,default=1" binding:"oneof=0 1"`
}

// gte 大于或等于
type UpdateTagRequest struct {
	ID         int `form:"id" binding:"required,gte=1"`
	Name       string `form:"name" binding:"min=3,max=100"`
	State      int    `form:"state,default=1" binding:"oneof=0 1"`
	ModifiedBy string `form:"modified_by" binding:"required,min=3,max=100"`
}

type DeleteTagRequest struct {
	ID int `form:"id" binding:"required,gte=1"`
}
