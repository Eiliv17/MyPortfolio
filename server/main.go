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

	api.GET("/posts", controllers.GetPosts)

	api.GET("/posts/:id", controllers.GetPost)

	api.GET("/projects", controllers.GetProjects)

	api.GET("/projects/:id", controllers.GetProject)

	r.Run() // listen and serve on 0.0.0.0:8080
}
