package router

import "github.com/gin-gonic/gin"

type PublicRouter struct {
}

func NewPublicRouter() IRouter {
	return &PublicRouter{}
}

func (p *PublicRouter) Register(group *gin.RouterGroup) {
	g := group.Group("/public")
	{
		g.GET("/ping1", func(ctx *gin.Context) {
			ctx.JSON(201, "ping1")
		})
	}
}
