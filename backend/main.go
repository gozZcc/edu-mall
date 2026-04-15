package main

import (
	"backend/config"
)

func main() {
	_, err := config.Init()
	if err != nil {
		panic("基础配置启动失败")
	}
}
