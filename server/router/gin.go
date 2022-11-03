package router

import "github.com/gin-gonic/gin"

func NewRouterGin(port string) Router {
	app := gin.New()
	app.Use(gin.Logger(), gin.Recovery())
	return &router{
		port: port,
		gin:  app,
	}
}

func (r *router) Start() {
	r.gin.Run(r.port)
}
