package router

import (
	"backend/controller"

	"github.com/gin-gonic/gin"
)

type PublicRouter struct {
	publicController controller.IPublicController
}

func NewPublicRouter(publicCtrl controller.IPublicController) IRouter {
	return &PublicRouter{
		publicController: publicCtrl,
	}
}

func (p *PublicRouter) Register(group *gin.RouterGroup) {
	g := group.Group("/public")
	{
		g.GET("/login", p.publicController.Login)
	}
}
