package controllers

import (
	"net/http"
	"portfolio-server/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetProjects(c *gin.Context) {

	page := c.Query("page")
	if page == "" {
		latestprojects, err := models.GetShortProjects(0, 10)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "couldn't get any project",
			})
			return
		}
		if latestprojects == nil {
			c.JSON(http.StatusBadRequest, gin.H{})
			return
		}

		c.IndentedJSON(http.StatusOK, latestprojects)
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

	pageprojects, err := models.GetShortProjects(offset, 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "couldn't get any project",
		})
		return
	}
	if pageprojects == nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	c.IndentedJSON(http.StatusOK, pageprojects)
}

func GetProject(c *gin.Context) {

	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id parameter",
		})
		return
	}

	project, err := models.GetProjectByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id parameter",
		})
		return
	}

	c.IndentedJSON(http.StatusOK, project)
}
