package server

import (
	"github.com/tgo-team/FeiGeIMServer/pkg/lmhttp"
	"github.com/tgo-team/FeiGeIMServer/pkg/log"
)

// Server server
type Server struct {
	r *lmhttp.LMHttp
	log.TLog
}

// New 创建server
func New() *Server {
	r := lmhttp.New()
	r.Use(lmhttp.CORSMiddleware())
	s := &Server{
		r: r,
	}
	return s
}

// Run 运行
func (s *Server) Run(addr ...string) error {
	s.r.Static("/swagger", "./configs/swagger")
	return s.r.Run(addr...)
}

// GetRoute 获取路由
func (s *Server) GetRoute() *lmhttp.LMHttp {
	return s.r
}
