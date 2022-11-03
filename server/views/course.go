package views

import (
	"simple-go-search/server/models"

	"github.com/google/uuid"
)

type CourseCreated struct {
	ID    uuid.UUID `json:"id"`
	Title string    `json:"title"`
}

func NewCourseCreated(courseModel *models.Course) *CourseCreated {
	return &CourseCreated{
		ID:    courseModel.ID,
		Title: courseModel.Title,
	}
}

type CourseFind struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	ImgCover    string    `json:"img_cover"`
	Description string    `json:"desc"`
	Price       int       `json:"price"`
	Category    string    `json:"category"`
}

func NewCoursesFind(coursesModel *[]models.Course) *[]CourseFind {

	var courses []CourseFind

	for _, course := range *coursesModel {
		var tempCourse = CourseFind{
			ID:          course.ID,
			Title:       course.Title,
			ImgCover:    course.ImgCover,
			Description: course.Description,
			Price:       course.Price,
			Category:    course.Category.String(),
		}
		courses = append(courses, tempCourse)
	}
	return &courses

}
