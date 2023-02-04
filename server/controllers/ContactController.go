package controllers

import (
	"net/http"
	"os"
	"portfolio-server/models"

	"github.com/gin-gonic/gin"
)

func OptionsContact(c *gin.Context) {
	if os.Getenv("GO_ENV") != "production" {
		c.Header("Access-Control-Allow-Origin", "*")
	}
	c.Header("Access-Control-Allow-Methods", "POST")
	c.Header("Access-Control-Allow-Headers", "content-type")
	c.AbortWithStatus(http.StatusNoContent)
}

func PostContact(c *gin.Context) {
	if os.Getenv("GO_ENV") != "production" {
		c.Header("Access-Control-Allow-Origin", "*")
	}

	c.Header("Access-Control-Allow-Methods", "POST")
	c.Header("Access-Control-Allow-Headers", "content-type")

	// body data struct
	body := struct {
		FullName string `json:"fullName" binding:"required"`
		Email    string `json:"email" binding:"email,required"`
		Subject  string `json:"subject" binding:"required"`
		Message  string `json:"message" binding:"required"`
	}{}

	// gets the body data
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.IndentedJSON(http.StatusOK, gin.H{
			"error": "all fields are required, and the email has to be an acutal email",
		})
		return
	}

	submission := models.CreateContactSubmission(
		body.FullName,
		body.Email,
		body.Subject,
		body.Message,
	)

	err = submission.Upload()
	if err != nil {
		c.IndentedJSON(http.StatusOK, gin.H{
			"error": "failed uploading your form submission",
		})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"result": "form submitted",
	})
}
