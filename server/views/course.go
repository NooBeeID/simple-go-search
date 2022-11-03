package views

import (
	"simple-go-search/server/models"

	"github.com/google/uuid"
)

type CourseCreated struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	ImgCover    string    `json:"img_cover"`
	Description string    `json:"desc"`
	Price       int       `json:"price"`
	Category    string    `json:"category"`
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
