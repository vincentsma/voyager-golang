package services

import (
	"github.com/gin-gonic/gin"
	"voyager-golang/models"
	"time"
	"strconv"
)

func AddPost(c *gin.Context) {
	var post models.Post
	c.Bind(&post)

	ret := models.AddPost(post)

	if ret == nil {
		c.JSON(201, post)
	} else {
		c.JSON(400, gin.H{"error": "bad post input"})
	}
}


func GetPosts(c *gin.Context) {
	//var posts []models.User

	posts, err := models.GetPosts()

	if err == nil {
		c.JSON(200, posts)
	} else {
		c.JSON(404, gin.H{"error": "no post(s) in the table"})
	}
}

func GetPost(c *gin.Context) {
	sid := c.Params.ByName("id")
	id, err := strconv.ParseInt(sid, 10, 64)

	if err != nil {
		panic(err)
	}

	post, err := models.GetPost(id)

	if err == nil {
		c.JSON(200, post)
	} else {
		c.JSON(404, gin.H{"error": "no post in the table"})
	}
}

func UpdatePost(c *gin.Context) {
	var post models.Post
	c.Bind(&post)

	sid := c.Params.ByName("id")
	id, err := strconv.ParseInt(sid, 10, 64)

	if err != nil {
		panic(err)
	}

	post.Updated = time.Now()

	ret := models.UpdatePost(id, post)

	if ret == nil {
		c.JSON(201, post)
	} else {
		c.JSON(400, gin.H{"error": "bad post input"})
	}
}