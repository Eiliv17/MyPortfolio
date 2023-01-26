package controllers

import (
	"net/http"
	"portfolio-server/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPosts(c *gin.Context) {

	page := c.Query("page")
	if page == "" {
		latestposts, err := models.GetShortPosts(0, 10)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "couldn't get any post",
			})
			return
		}
		if latestposts == nil {
			c.JSON(http.StatusBadRequest, gin.H{})
			return
		}

		c.IndentedJSON(http.StatusOK, latestposts)
		return
	}

	pageNumber, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid page",
		})
		return
	}

	var offset int64 = pageNumber*10 - 9

	pageposts, err := models.GetShortPosts(offset, 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "couldn't get any post",
		})
		return
	}
	if pageposts == nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	c.IndentedJSON(http.StatusOK, pageposts)
}

func GetPost(c *gin.Context) {

	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id parameter",
		})
		return
	}

	post, err := models.GetPostByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id parameter",
		})
		return
	}

	c.IndentedJSON(http.StatusOK, post)
}
