package gormRepo

import (
	"errors"
	"simple-go-search/db"
	"simple-go-search/server/models"
	"simple-go-search/server/repositories"
	"simple-go-search/server/views"
	"strings"

	"gorm.io/gorm"
)

type gormDB struct {
	db *gorm.DB
}

func NewGormCourse(db *db.DB) repositories.CourseRepository {
	return &gormDB{
		db: db.GormDB,
	}
}

func (c *gormDB) Create(course *models.Course) error {
	return c.db.Create(course).Error
}

func (c *gormDB) FindCourseByFilter(filter *models.CourseFilter) (*[]models.Course, error) {
	var courses []models.Course

	title := "%" + strings.ToLower(filter.Title) + "%"
	desc := "%" + strings.ToLower(filter.Description) + "%"
	category := "%" + strings.ToLower(filter.Category.String()) + "%"

	rows := c.db.Where(`
			LOWER(title) LIKE ? 
			AND LOWER(description) LIKE ? 
			AND LOWER(category) LIKE ?
		`, title, desc, category).Find(&courses)

	if rows.Error != nil {
		return nil, rows.Error
	}

	if rows.RowsAffected == 0 {
		return nil, errors.New(string(views.Err_NotFound))
	}

	return &courses, nil
}
