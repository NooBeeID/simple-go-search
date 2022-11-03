package router

import (
	"simple-go-search/server/controllers"

	"github.com/gin-gonic/gin"
)

func NewRouterGin(port string, course controllers.CourseController) Router {
	app := gin.New()
	app.Use(gin.Logger(), gin.Recovery())
	return &router{
		port:   port,
		gin:    app,
		course: course,
	}
}

func (r *router) Start() {
	api := r.gin.Group("api")
	v1 := api.Group("v1")
	r.v1Route(v1)

	r.gin.Run(r.port)
}

func (r *router) v1Route(route *gin.RouterGroup) {
	course := route.Group("/courses")
	{
		course.POST("", r.course.CreateNewCourse)
	}
}
