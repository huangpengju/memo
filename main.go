package main

import (
	"memo/conf"
	"memo/routes"
)

func main() {
	// 初始配置
	conf.Init()

	r := routes.NewRouter()
	// 使用 3000 端口运行程序
	_ = r.Run(conf.HttpPort)
}
