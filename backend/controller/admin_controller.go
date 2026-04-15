package controller

import "github.com/gin-gonic/gin"

type AdminController struct {
}

func NewAdminController() *AdminController {
	return &AdminController{}
}

func (c *AdminController) Ping(ctx *gin.Context) {
	ctx.JSON(201, "admin ping")
}
