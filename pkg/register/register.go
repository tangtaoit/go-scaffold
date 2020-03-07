package register

import "github.com/tgo-team/FeiGeIMServer/pkg/lmhttp"

// APIRouter API路由
type APIRouter interface {
	Route(r *lmhttp.LMHttp)
}

var apiRoutes = make([]APIRouter, 0)

// Add 添加api路由
func Add(r APIRouter) {
	apiRoutes = append(apiRoutes, r)
}

// GetRoutes 获取所有api路由
func GetRoutes() []APIRouter {
	return apiRoutes
}
