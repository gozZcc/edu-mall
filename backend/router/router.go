package router

import "github.com/gin-gonic/gin"

// 路由

type IRouter interface {
	Register(engine *gin.RouterGroup)
}

// 全局路由前缀常量（唯一修改点，无硬编码）
const ApiPrefix = "/api/v1"

func Init() *gin.Engine {
	ge := gin.Default()
	// 统一前缀：从常量读取，不写死
	rg := ge.Group(ApiPrefix)
	// 注册所有模块路由
	routers := []IRouter{
		NewPublicRouter(),
		NewCostomerRouter(),
		NewAdminRouter(),
	}
	// 统一注册
	for _, r := range routers {
		r.Register(rg)
	}
	return ge
}
