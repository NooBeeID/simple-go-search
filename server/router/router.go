package router

import (
	"simple-go-search/server/controllers"

	"github.com/gin-gonic/gin"
)

type router struct {
	port string
	gin  *gin.Engine

	course controllers.CourseController
}

type Router interface {
	Start()
}
