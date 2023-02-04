package controllers

import (
	"net/http"
	"os"
	"portfolio-server/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RetrieveProjects(c *gin.Context) {
	if os.Getenv("GO_ENV") != "production" {
		c.Header("Access-Control-Allow-Origin", "*")
	}

	reqoffset := c.Query("offset")
	reqlimit := c.Query("limit")

	// set default values if empty
	if reqoffset == "" {
		reqoffset = "0"
	}

	if reqlimit == "" {
		reqlimit = "10"
	}

	// convert data type
	offset, err := strconv.ParseInt(reqoffset, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid offset",
		})
		return
	}

	limit, err := strconv.ParseInt(reqlimit, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid limit",
		})
		return
	}

	// retrieve the projects and check for any errors
	projects, err := models.GetShortProjects(offset, limit)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "couldn't get any article",
		})
		return
	}
	if projects == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "couldn't get any article",
		})
		return
	}

	c.IndentedJSON(http.StatusOK, projects)
}

func RetrieveProject(c *gin.Context) {
	if os.Getenv("GO_ENV") != "production" {
		c.Header("Access-Control-Allow-Origin", "*")
	}

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
