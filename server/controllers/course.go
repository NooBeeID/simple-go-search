package controllers

import (
	"fmt"
	"simple-go-search/server/params"
	"simple-go-search/server/services"
	"simple-go-search/server/views"

	"github.com/gin-gonic/gin"
)

type CourseController interface {
	CreateNewCourse(c *gin.Context)
	FilterCourse(c *gin.Context)
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

func (course *courseController) FilterCourse(c *gin.Context) {
	var req params.CourseFilter
	req.Title = c.Query("title")
	req.Category = c.Query("category")
	req.Description = c.Query("description")

	fmt.Printf("%+v\n", req)

	resp := course.svc.FilterCourses(&req)
	WriteJsonResponse(c, resp)
}
