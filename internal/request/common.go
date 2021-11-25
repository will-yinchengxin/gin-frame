package request

type PageInfo struct {
	Page     int64 `form:"page" json:"page" validate:"omitempty,gt=0" label:"分页"`
	PageSize int64 `form:"pageSize" json:"pageSize" validate:"omitempty,gt=0" label:"每页条数"`
}

type Time struct {
	StartTime int64 `form:"startTime" json:"startTime" validate:"required,ltNowTime" label:"开始时间"`
	EndTime   int64 `form:"endTime" json:"endTime" validate:"required,unixTime,gtefield=StartTime" label:"截止时间"`
}
