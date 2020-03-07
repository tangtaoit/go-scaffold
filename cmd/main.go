package main

import (
	"github.com/tgo-team/FeiGeIMServer/internal/api"
	"github.com/tgo-team/FeiGeIMServer/internal/config"
	"github.com/tgo-team/FeiGeIMServer/internal/server"
)

func main() {
	cfg := config.New()
	// 创建server
	s := server.New()
	// 初始化context
	ctx := api.NewContext(cfg)
	// 初始化api
	api.Init(ctx)
	// 开始route
	api.Route(s.GetRoute())
	// 运行
	err := s.Run(":8080")
	if err != nil {
		panic(err)
	}
}
