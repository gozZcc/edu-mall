package main

import (
	"backend/config"
	"backend/controller"
	"backend/repository"
	"backend/router"
	"backend/service"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// 初始化logger

	// 获取配置信息
	c, err := config.Init()
	if err != nil {
		panic("获取配置信息失败")
	}
	// 连接数据库
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=true&loc=Local",
		c.Mysql.User, c.Mysql.Password, c.Mysql.Host, c.Mysql.Port, c.Mysql.Database, c.Mysql.Charset)))
	if err != nil {
		panic("连接数据库失败")
	}
	// 初始化repository层
	repos := repository.NewRepositories(db)
	// 初始化service层,注入repository
	svcs := service.NewServices(repos)
	// 初始化controller层,注入services
	ctrls := controller.NewControllers(svcs)
	// 初始化router层,注入controllers
	r := router.NewRouter(ctrls)

	// 启动服务
	err = r.Run(fmt.Sprintf(":%s", c.Server.HttpPort))
	if err != nil {
		panic("启动服务失败")
	}
}
