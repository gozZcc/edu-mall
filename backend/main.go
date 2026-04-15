package main

import (
	"backend/config"
	"backend/router"
	"fmt"
)

func main() {
	c, err := config.Init()
	if err != nil {
		panic("基础配置启动失败")
	}
	// 注册路由
	re := router.Init()
	// 启动服务
	re.Run(fmt.Sprintf(":%v", c.Server.HttpPort))
}
