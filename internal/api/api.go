package api

import (
	"github.com/tgo-team/FeiGeIMServer/pkg/lmhttp"
	"github.com/tgo-team/FeiGeIMServer/pkg/register"
)

// Route 路由
func Route(r *lmhttp.LMHttp) {
	routes := register.GetRoutes()
	for _, route := range routes {
		route.Route(r)
	}
}

// Init 注册所有api
func Init(ctx *Context) {

}
