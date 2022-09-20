package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mahiro72/go-mysql-ca/config"
)

type Router struct {
	engine *gin.Engine
}

// NewRouterは新しいRouterを初期化し構造体のポインタを返します
func NewRouter() *Router {
	e := gin.Default()
	return &Router{
		engine: e,
	}
}

// Serveはhttpサーバーを起動します
func (r *Router) Serve() {
	r.engine.Run(fmt.Sprintf(":%s", config.Port()))
}
