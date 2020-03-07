package lmhttp

import (
	"net/http"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/tgo-team/FeiGeIMServer/pkg/cache"
)

// LMHttp LMHttp
type LMHttp struct {
	r    *gin.Engine
	pool sync.Pool
}

// New New
func New() *LMHttp {
	l := &LMHttp{
		r:    gin.Default(),
		pool: sync.Pool{},
	}
	l.r.Use(gin.Recovery())
	l.pool.New = func() interface{} {
		return allocateContext()
	}
	return l
}

func allocateContext() *Context {
	return &Context{Context: nil}
}

// Context Context
type Context struct {
	*gin.Context
}

func (c *Context) reset() {
	c.Context = nil
}

// ResponseError ResponseError
func (c *Context) ResponseError(err error) {
	c.JSON(http.StatusBadRequest, gin.H{
		"msg":    err.Error(),
		"status": http.StatusBadRequest,
	})
}

// ResponseOK 返回成功
func (c *Context) ResponseOK() {
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
	})
}

// Response Response
func (c *Context) Response(data interface{}) {
	c.JSON(http.StatusOK, data)
}

// ResponseWithStatus ResponseWithStatus
func (c *Context) ResponseWithStatus(status int, data interface{}) {
	c.JSON(status, data)
}

// HandlerFunc HandlerFunc
type HandlerFunc func(c *Context)

// LMHttpHandler LMHttpHandler
func (l *LMHttp) LMHttpHandler(handlerFunc HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		//hc := l.pool.Get().(*Context)
		//hc.reset()
		//hc.Context = c
		//handlerFunc(hc)
		//l.pool.Put(hc)
		handlerFunc(&Context{Context: c})
	}
}

// Run Run
func (l *LMHttp) Run(addr ...string) error {
	return l.r.Run(addr...)
}

// POST POST
func (l *LMHttp) POST(relativePath string, handlers ...HandlerFunc) {
	l.r.POST(relativePath, l.handlersToGinHandleFuncs(handlers)...)
}

// GET GET
func (l *LMHttp) GET(relativePath string, handlers ...HandlerFunc) {
	l.r.GET(relativePath, l.handlersToGinHandleFuncs(handlers)...)
}

// Static Static
func (l *LMHttp) Static(relativePath string, root string) {
	l.r.Static(relativePath, root)
}

// Use Use
func (l *LMHttp) Use(handlers ...HandlerFunc) {
	l.r.Use(l.handlersToGinHandleFuncs(handlers)...)
}

// ServeHTTP ServeHTTP
func (l *LMHttp) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	l.r.ServeHTTP(w, req)
}

// Group Group
func (l *LMHttp) Group(relativePath string, handlers ...HandlerFunc) *RouterGroup {
	return newRouterGroup(l.r.Group(relativePath, l.handlersToGinHandleFuncs(handlers)...), l)
}

func (l *LMHttp) handlersToGinHandleFuncs(handlers []HandlerFunc) []gin.HandlerFunc {
	newHandlers := make([]gin.HandlerFunc, 0, len(handlers))
	if handlers != nil {
		for _, handler := range handlers {
			newHandlers = append(newHandlers, l.LMHttpHandler(handler))
		}
	}
	return newHandlers
}

// AuthMiddleware 认证中间件
func (l *LMHttp) AuthMiddleware(cache cache.Cache, tokenPrefix string) HandlerFunc {

	return func(c *Context) {
		token := c.GetHeader("token")
		uidAndName := GetLoginUID(token, tokenPrefix, cache)
		if strings.TrimSpace(uidAndName) == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"msg": "请先登录！",
			})
			return
		}
		uidAndNames := strings.Split(uidAndName, "@")
		if len(uidAndNames) != 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"msg": "token有误！",
			})
			return
		}
		c.Set("uid", uidAndNames[0])
		c.Set("name", uidAndNames[1])
		c.Next()
	}
}

// GetLoginUID GetLoginUID
func GetLoginUID(token string, tokenPrefix string, cache cache.Cache) string {
	uid, err := cache.Get(tokenPrefix + token)
	if err != nil {
		return ""
	}
	return uid
}

// RouterGroup RouterGroup
type RouterGroup struct {
	*gin.RouterGroup
	L *LMHttp
}

func newRouterGroup(g *gin.RouterGroup, l *LMHttp) *RouterGroup {
	return &RouterGroup{RouterGroup: g, L: l}
}

// POST POST
func (r *RouterGroup) POST(relativePath string, handlers ...HandlerFunc) {
	r.RouterGroup.POST(relativePath, r.L.handlersToGinHandleFuncs(handlers)...)
}

// GET GET
func (r *RouterGroup) GET(relativePath string, handlers ...HandlerFunc) {
	r.RouterGroup.GET(relativePath, r.L.handlersToGinHandleFuncs(handlers)...)
}

// DELETE DELETE
func (r *RouterGroup) DELETE(relativePath string, handlers ...HandlerFunc) {
	r.RouterGroup.DELETE(relativePath, r.L.handlersToGinHandleFuncs(handlers)...)
}

// PUT PUT
func (r *RouterGroup) PUT(relativePath string, handlers ...HandlerFunc) {
	r.RouterGroup.PUT(relativePath, r.L.handlersToGinHandleFuncs(handlers)...)
}

//CORSMiddleware 跨域
func CORSMiddleware() HandlerFunc {

	return func(c *Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, token, accept, origin, Cache-Control, X-Requested-With, app_id, open_id, noncestr, sign, timestamp,store_token")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT,DELETE,PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
