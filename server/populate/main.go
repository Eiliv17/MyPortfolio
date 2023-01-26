package main

import (
	"portfolio-server/initializers"
	"portfolio-server/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.LoadDatabase()
}

func main() {

	for i := 0; i < 10; i++ {
		post := models.CreatePost("Hello", "Hello this is a description", "this is the body text", []string{"test", "testhello"})
		post.Upload()
	}

}
