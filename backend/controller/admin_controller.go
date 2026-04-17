package controller

import (
	"backend/consts"
	dto "backend/dto/admin"
	"backend/service"

	"github.com/gin-gonic/gin"
)

type AdminController struct {
	adminService *service.AdminService
}

func NewAdminController(services *service.Services) *AdminController {
	return &AdminController{
		adminService: services.AdminService,
	}
}

func (c *AdminController) GetInfo(ctx *gin.Context) {
	// 解析参数
	getInfoReq := &dto.AdminGetInfoReq{}
	ctx.ShouldBindJSON(&getInfoReq)
	adminUserBO, err := c.adminService.GetAdminUserInfoByID(ctx.Request.Context(), getInfoReq.ID)
	if err != nil {
		Failed(ctx, consts.CodeServerInternal)
		return
	}
	Success(ctx, dto.AdminGetInfoResp{
		Name: adminUserBO.Name,
	})
}
