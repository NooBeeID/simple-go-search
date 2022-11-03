package gormRepo

import (
	"simple-go-search/db"
	"simple-go-search/server/models"
	"simple-go-search/server/repositories"

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

	title := "%" + filter.Title + "%"
	desc := "%" + filter.Description + "%"
	category := "%" + filter.Category + "%"
	price := filter.Price

	err := c.db.Where(`
			title LIKE ? 
			OR description LIKE ? 
			OR price=? 
			OR category LIKE ?
		`, title, desc, price, category).Find(&courses).Error

	if err != nil {
		return nil, err
	}

	return &courses, nil
}
