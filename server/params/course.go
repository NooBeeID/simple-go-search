package params

import "simple-go-search/server/models"

type CreateCourse struct {
	Title       string `json:"title"`
	ImgCover    string `json:"img_cover"`
	Description string `json:"desc"`
	Price       int    `json:"price"`
	Category    string `json:"category"`
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
