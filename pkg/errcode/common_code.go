package errcode

var (
	Success          = NewError(200, "成功")
	ServerError      = NewError(9000, "服务内部错误")
	InvalidParameter = NewError(5003, "无效参数")

	TokenTypeError          = NewError(5020, "Token类型不正确")
	TokenInvalid            = NewError(5021, "无效的用户令牌")
	AuthServerTokenExpire   = NewError(5022, "认证中心token过期，请重新登录")
	AuthServerGetDataFailed = NewError(5023, "从认证中心获取数据获取失败，请重试")
	TokenCreateError        = NewError(5024, "登陆凭证获取失败，请重试")
	ThirdTokenBindError     = NewError(5025, "第三方平台绑定token不合法")

	// Tag Error
	ErrorGetTagListFail = NewError(200010, "获取标签列表失败")
	ErrorCountTagFail   = NewError(200011, "统计标签数量失败")
	ErrorCreateTagFail  = NewError(200012, "创建标签失败")
	ErrorUpdateTagFail  = NewError(200013, "更新标签失败")
	ErrorDeleteTagFail  = NewError(200014, "删除标签失败")
)
