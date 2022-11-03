package controllers

import (
	"simple-go-search/server/views"

	"github.com/gin-gonic/gin"
)

func WriteJsonResponse(c *gin.Context, payload *views.Response) {
	c.JSON(payload.Status, payload)
}
