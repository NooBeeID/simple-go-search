package services

import (
	"simple-go-search/server/params"
	"simple-go-search/server/repositories"
	"simple-go-search/server/views"
)

type CourseServices interface {
	CreateNewCourse(req *params.CreateCourse) *views.Response
}

type courseSvc struct {
	repo repositories.CourseRepository
}

func NewCourseSvc(repo repositories.CourseRepository) CourseServices {
	return &courseSvc{
		repo: repo,
	}
}

func (c *courseSvc) CreateNewCourse(req *params.CreateCourse) *views.Response {
	course := req.ParseToModel()

	err := c.repo.Create(course)
	if err != nil {
		return views.RepositoryError(err)
	}

	courseView := views.NewCourseCreated(course)
	return views.SuccessCreated(courseView)
}
