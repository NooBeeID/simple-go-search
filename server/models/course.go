package models

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Course struct {
	ID          uuid.UUID
	Title       string
	ImgCover    string
	Description string
	Price       int
	Category    Category
}

type CourseFilter struct {
	Title       string
	Description string
	Price       int
	Category    Category
}

type Category string

const (
	Category_Frontend Category = "Frontend"
	Category_Backend  Category = "Backend"
	Category_UIUX     Category = "UIUX"
	Category_DevOps   Category = "DevOps"
)

func NewCategory(cat string) Category {
	return Category(cat)
}

func (c Category) Validate() (err error) {
	switch c {
	case Category_Frontend:
		return
	case Category_Backend:
		return
	case Category_DevOps:
		return
	case Category_UIUX:
		return
	default:
		err = fmt.Errorf("category `%v` not found", c)
		return
	}
}

func (c Category) String() string {
	return string(c)
}

func (c *Course) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.New()
	err = c.Category.Validate()
	return

}
