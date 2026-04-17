package router

const (
	ApiPrefix     = "/api"                       // 根接口前缀
	ApiVersion    = "v1"                         // 版本号
	FullApiPrefix = ApiPrefix + "/" + ApiVersion // /api/v1

	// 模块路由
	ModulePublic = "/public"
	ModuleAdmin  = "/admin"
)
