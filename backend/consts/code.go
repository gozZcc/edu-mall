package consts

// 企业级错误码分段规范（互联网公司通用标准）
// 0         : 业务处理成功，固定值，全局唯一
// 1xxxx     : 通用基础错误（参数、权限、数据转换等）
// 2xxxx     : 用户模块业务错误（登录、注册、用户信息等）
// 3xxxx~4xxxx: 其他业务模块错误（按业务线扩展）
// 5xxxx     : 服务端内部错误（兜底异常）
const (
	CodeSuccess        = 0     // 业务成功
	CodeParamError     = 10001 // 参数校验/解析失败
	CodeConvError      = 10002 // 数据转换失败
	CodeUnauthorized   = 10003 // 未授权/未登录
	CodeForbidden      = 10004 // 权限不足
	CodeLoginFailed    = 20001 // 登录业务失败
	CodeServerInternal = 50000 // 服务端内部错误
)
