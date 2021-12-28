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

type OnlineList struct {
	//StartTime int64 `form:"startTime" json:"startTime" validate:"required,unixTime" label:"开始时间"`
	//EndTime   int64 `form:"endTime" json:"endTime" validate:"required,unixTime,gtfield=StartTime" label:"截止时间"`

	Type       int64  `json:"type" validate:"required,oneof=1 2"`
	Day        int64  `json:"day" validate:"required_if=Type 2,gte=0,lte=365"`  // 字段关联输入, required 和 required_if
	
	ChannelId string `form:"channelId" json:"channelId" validate:"omitempty" label:"渠道id"`
	AppId     int64  `form:"appId" json:"appId" validate:"omitempty,gt=0" label:"应用id"`
	Url       string `form:"url" json:"url" validate:"omitempty,urlFormat"` // 使用自定义校验规则
}

type Company struct {
	ID int64 `json:"id" validate:"required,gt=0"`
}
