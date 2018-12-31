package handler

import (
	"github.com/sundogrd/content-api/services/content"
	"github.com/gin-gonic/gin"
)

// GetContent ...
func GetContent(c *gin.Context) {
	content.ContentRepositoryInstance().FindOne(c, content.FindOneRequest{ID: 1})
	c.JSON(200, gin.H{
		"msg": "asd",
	})
}

// ListContent ...
func ListContent(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "list",
	})
}

func CreateContent(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "create",
	})
}

func UpdateContent(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "updated",
	})
}

func DeleteContent(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "deleted",
	})
}
