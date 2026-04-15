package api

// LoginReq 登录接口请求
type LoginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginResp 登录接口响应
type LoginResp struct {
	Token    string `json:"token"`
	Username string `json:"username"`
}
