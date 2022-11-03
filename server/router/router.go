package router

import "github.com/gin-gonic/gin"

type router struct {
	port string
	gin  *gin.Engine
}

type Router interface {
	Start()
}
