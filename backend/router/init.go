package router

import (
	"backend/controller"

	"github.com/gin-gonic/gin"
)

func NewRouter(controllers *controller.Controllers) *gin.Engine {
	r := gin.Default()
	v1Group := r.Group(FullApiPrefix)
	{
		RegisterPublicRouter(v1Group, controllers)
	}
	return r
}
