package router

import (
	"backend/controller"

	"github.com/gin-gonic/gin"
)

type AdminRouter struct {
}

func NewAdminRouter() *AdminRouter {
	return &AdminRouter{}
}

func RegisterAdminRouter(r *gin.RouterGroup, ctrls *controller.Controllers) {
	pr := r.Group(ModuleAdmin)
	{
		pr.POST("/getInfo", ctrls.AdminController.GetInfo)
	}

}
