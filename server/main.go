package main

import (
	"portfolio-server/controllers"
	"portfolio-server/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.LoadDatabase()
}

func main() {
	r := gin.Default()

	api := r.Group("/api")

	api.GET("/posts", controllers.RetrievePosts)
	api.GET("/posts/:id", controllers.RetrievePost)

	api.GET("/projects", controllers.RetrieveProjects)
	api.GET("/projects/:id", controllers.RetrieveProject)

	api.POST("/contact", controllers.PostContact)
	api.OPTIONS("/contact", controllers.OptionsContact)

	r.Run() // listen and serve on 0.0.0.0:8080
}
