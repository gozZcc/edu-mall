package router

import (
	"backend/controller"

	"github.com/gin-gonic/gin"
)

type CostomerRouter struct {
}

func NewCostomerRouter() IRouter {
	return &CostomerRouter{}
}

func (c *CostomerRouter) Register(group *gin.RouterGroup) {
	g := group.Group("/customer")
	ctrl := controller.NewCustomerController()
	{
		g.GET("/ping", ctrl.Ping)
	}
}
