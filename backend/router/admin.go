package router

import (
	"backend/controller"

	"github.com/gin-gonic/gin"
)

type AdminRouter struct {
}

func NewAdminRouter() IRouter {
	return &AdminRouter{}
}

func (a *AdminRouter) Register(group *gin.RouterGroup) {
	g := group.Group("/admin")
	ctrl := controller.NewAdminController()
	{
		g.GET("/ping", ctrl.Ping)
	}
}
