package consts

// 错误码对应前端友好提示，与code.go严格一一对应，禁止出现码值无对应提示
const (
	MsgSuccess        = "请求成功"
	MsgParamError     = "请求参数格式错误，请检查后重试"
	MsgConvError      = "数据处理失败"
	MsgUnauthorized   = "请先登录"
	MsgForbidden      = "权限不足，无法操作"
	MsgLoginFailed    = "账号或密码错误"
	MsgServerInternal = "服务器内部错误，请稍后重试"
)

// ErrMsgMap 全局错误码-提示映射，唯一可信来源，杜绝业务代码硬编码
var ErrMsgMap = map[int]string{
	CodeSuccess:        MsgSuccess,
	CodeParamError:     MsgParamError,
	CodeConvError:      MsgConvError,
	CodeUnauthorized:   MsgUnauthorized,
	CodeForbidden:      MsgForbidden,
	CodeLoginFailed:    MsgLoginFailed,
	CodeServerInternal: MsgServerInternal,
}

// GetErrMsg 根据错误码获取对应提示，兜底返回服务端错误，避免空提示
func GetErrMsg(code int) string {
	if msg, ok := ErrMsgMap[code]; ok {
		return msg
	}
	return MsgServerInternal
}
