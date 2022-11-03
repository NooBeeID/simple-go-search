package controllers

import (
	"simple-go-search/server/params"
	"simple-go-search/server/services"
	"simple-go-search/server/views"

	"github.com/gin-gonic/gin"
)

type CourseController interface {
	CreateNewCourse(c *gin.Context)
}

type courseController struct {
	svc services.CourseServices
}

func NewCourseController(svc services.CourseServices) CourseController {
	return &courseController{
		svc: svc,
	}
}

func (course *courseController) CreateNewCourse(c *gin.Context) {
	var req params.CreateCourse

	err := c.ShouldBindJSON(&req)
	if err != nil {
		WriteJsonResponse(c, views.BadRequest(err))
		return
	}

	resp := course.svc.CreateNewCourse(&req)
	WriteJsonResponse(c, resp)
}
