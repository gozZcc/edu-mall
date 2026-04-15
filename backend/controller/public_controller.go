package controller

import (
	"backend/api"
	"backend/consts"
	"backend/service"

	"github.com/gin-gonic/gin"
)

type IPublicController interface {
	Login(ctx *gin.Context)
}

type PublicController struct {
	userService service.IUserService
}

func NewPublicController(userService service.IUserService) IPublicController {
	return &PublicController{
		userService: userService,
	}
}

type GetPublicPingReq struct{}

type GetPublicPingResp struct {
	Message string `json:"message"` // 业务字段明确定义
}

func (c *PublicController) Login(ctx *gin.Context) {
	// 1. 绑定前端参数到 api.UserLoginReq
	var req api.LoginReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		Failed(ctx, consts.CodeParamError, err.Error())
		return
	}

	// serviceInput :=

	// 2. 调用 Service，获取 Entity

	// _ = service.UserService.Login(req.Username, req.Password)
	// // if err != nil {
	// // 	Failed(ctx, consts.CodeFailed, consts.MsgFailed)
	// // 	return
	// // }

	// // 3. 生成 Token
	// token := "mock_token_" + dbUser.Username

	// // 4. 【核心】copier 转换：Entity → Resp
	// var resp api.LoginResp
	// resp.Token = token
	// if err := copier.Copy(&resp, dbUser); err != nil {
	// 	Failed(ctx, consts.CodeFailed, consts.MsgConvError)
	// 	return
	// }
	// 补充自定义字段
	// 3. 返回 api.UserLoginResp
	// Success(ctx, resp)
}
