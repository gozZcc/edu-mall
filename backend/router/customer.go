package router

import "github.com/gin-gonic/gin"

type CostomerRouter struct {
}

func NewCostomerRouter() IRouter {
	return &CostomerRouter{}
}

func (c *CostomerRouter) Register(group *gin.RouterGroup) {
	g := group.Group("/customer")
	{
		g.GET("/ping2", func(ctx *gin.Context) {
			ctx.JSON(201, "ping2")
		})
	}
}
