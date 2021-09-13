package errcode

var (
	Success          = NewError(200, "成功")
	ServerError      = NewError(9000, "服务内部错误")
	FAILURE          = NewError(9001, "请求失败")
	GlobalError      = NewError(9002, "服务器异常,请稍后重试")
	Misinformation   = NewError(4100, "传参有误")
	InvalidParameter = NewError(5003, "无效参数")

	TokenTypeError          = NewError(5020, "Token类型不正确")
	TokenInvalid            = NewError(5021, "无效的用户令牌")
	AuthServerTokenExpire   = NewError(5022, "认证中心token过期，请重新登录")
	AuthServerGetDataFailed = NewError(5023, "从认证中心获取数据获取失败，请重试")
	TokenCreateError        = NewError(5024, "登陆凭证获取失败，请重试")
	ThirdTokenBindError     = NewError(5025, "第三方平台绑定token不合法")

	// db
	DBOperatorError = NewError(5026, "数据操作错误，请求失败")
	DBCreateError   = NewError(5027, "数据创建出错，请求失败")
	DBUpdateError   = NewError(5028, "数据更新出错，请求失败")
	DBDeleteError   = NewError(5029, "数据删除出错，请求失败")
)
