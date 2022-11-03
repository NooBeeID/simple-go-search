package params

import "simple-go-search/server/models"

type CreateCourse struct {
	Title       string
	ImgCover    string
	Description string
	Price       int
	Category    string
}

func (c *CreateCourse) ParseToModel() *models.Course {
	category := models.Category(c.Category)
	return &models.Course{
		Title:       c.Title,
		ImgCover:    c.ImgCover,
		Description: c.Description,
		Price:       c.Price,
		Category:    category,
	}
}
