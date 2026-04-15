package router

import "github.com/gin-gonic/gin"

type AdminRouter struct {
}

func NewAdminRouter() IRouter {
	return &AdminRouter{}
}

func (a *AdminRouter) Register(group *gin.RouterGroup) {
	g := group.Group("/admin")
	{
		g.GET("/ping3", func(ctx *gin.Context) {
			ctx.JSON(201, "ping3")
		})
	}
}
