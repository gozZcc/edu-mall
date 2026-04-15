package router

import (
	"backend/controller"
	"backend/service"

	"github.com/gin-gonic/gin"
)

// 路由

type IRouter interface {
	Register(engine *gin.RouterGroup)
}

// 全局路由前缀常量（唯一修改点，无硬编码）
const ApiPrefix = "/api/v1"

func Init() *gin.Engine {
	ge := gin.Default()

	// 1. 先初始化 Service
	userService := service.NewUserService()

	// 2. 初始化 Controller，只注入自己需要的 Service
	publicCtl := controller.NewPublicController(userService)
	// customerCtrl := controller.NewCustomerController()
	// adminCtrl := controller.NewAdminController()

	// 3.初始化 router 注入 controller
	publicRouter := NewPublicRouter(publicCtl)
	// 4. 初始化 Router 注入controller
	// publicRouter := router.NewPublicRouter(publicCtl)
	// adminRouter := router.NewAdminRouter(adminCtl)

	// 统一前缀：从常量读取，不写死
	rg := ge.Group(ApiPrefix)
	// 注册所有模块路由
	routers := []IRouter{
		// NewPublicRouter(),
		// NewCostomerRouter(),
		// NewAdminRouter(),
		publicRouter,
	}
	// 统一注册
	for _, r := range routers {
		r.Register(rg)
	}
	return ge
}
