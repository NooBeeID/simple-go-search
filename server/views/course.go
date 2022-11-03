package views

import (
	"simple-go-search/server/models"

	"github.com/google/uuid"
)

type CourseCreated struct {
	ID          uuid.UUID
	Title       string
	ImgCover    string
	Description string
	Price       int
	Category    string
}

func NewCourseCreated(courseModel *models.Course) *CourseCreated {
	return &CourseCreated{
		ID:          courseModel.ID,
		Title:       courseModel.Title,
		ImgCover:    courseModel.ImgCover,
		Description: courseModel.Description,
		Price:       courseModel.Price,
		Category:    courseModel.Category.String(),
	}
}
