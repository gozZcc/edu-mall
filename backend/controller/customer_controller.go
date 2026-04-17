package controller

import "github.com/gin-gonic/gin"

type CustomerController struct {
}

func NewCustomerController() *CustomerController {
	return &CustomerController{}
}

func (c *CustomerController) Ping(ctx *gin.Context) {
	ctx.JSON(201, "customer ping")
}
