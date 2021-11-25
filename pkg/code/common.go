package code

var (
	SUCCESS        = &CodeType{200, "请求成功"}
	ParamsError    = &CodeType{5101, "参数错误"}
	RecordNotFound = &CodeType{5102, "记录不存在"}
	ReportError    = &CodeType{5103, "数据上报失败"}
	MqError        = &CodeType{5104, "mq错误"}
)
