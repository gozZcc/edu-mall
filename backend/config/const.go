package config

const (
	// 配置文件相关
	ConfigDir     = "config/" // 配置文件目录
	DefaultEnv    = "dev"     // 默认环境
	ConfigFileExt = ".yaml"   // 配置文件后缀

	// 环境变量名
	EnvKeyAppEnv = "APP_ENV"     // 环境标识变量名
	EnvKeyDBPwd  = "DB_PASSWORD" // 数据库密码变量名（生产环境用）

	// 环境标识值
	EnvProd = "prod" // 生产环境标识
)
