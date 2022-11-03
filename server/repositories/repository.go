package repositories

import (
	"simple-go-search/server/models"
)

type CourseRepository interface {
	Create(course *models.Course) error
	FindCourseByFilter(filter *models.CourseFilter) (*[]models.Course, error)
}
