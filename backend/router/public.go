package router

import (
	"backend/controller"

	"github.com/gin-gonic/gin"
)

type PublicRouter struct {
}

func NewPublicRouter() *PublicRouter {
	return &PublicRouter{}
}

func RegisterPublicRouter(r *gin.RouterGroup, ctrls *controller.Controllers) {
	pr := r.Group(ModulePublic)
	{
		pr.GET("hello", ctrls.PublicController.Hello)
	}

}
