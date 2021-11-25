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

	ChannelId string `form:"channelId" json:"channelId" validate:"omitempty" label:"渠道id"`
	AppId     int64  `form:"appId" json:"appId" validate:"omitempty,gt=0" label:"应用id"`
	Url       string `form:"url" json:"url" validate:"omitempty,urlFormat"` // 使用自定义校验规则
}
