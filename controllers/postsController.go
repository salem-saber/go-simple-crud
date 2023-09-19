package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/salemsaber/03_crud/initializers"
	"github.com/salemsaber/03_crud/models"
)

func PostsCreate(c *gin.Context) {

	var body struct {
		Title string
		Body  string
	}
	c.Bind(&body)

	post := models.Post{
		Title: body.Title,
		Body:  body.Body,
	}

	result := initializers.DB.Create(&post) // pass pointer of data to Create

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})

}

func PostsIndex(c *gin.Context) {

	var posts []models.Post
	initializers.DB.Find(&posts)

	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {

	id := c.Param("id")
	var post models.Post
	initializers.DB.First(&post, id)

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context) {

	id := c.Param("id")
	var body struct {
		Title string
		Body  string
	}
	c.Bind(&body)

	var post models.Post

	initializers.DB.First(&post, id)

	initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body})

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsDelete(c *gin.Context) {

	id := c.Param("id")
	var post models.Post

	initializers.DB.Delete(&post, id)

	c.JSON(200, gin.H{})
}
