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

	/* for i := 0; i < 10; i++ {
		post := models.CreatePost("Hello", "Hello this is a description", "this is the body text", []string{"test", "testhello"})
		post.Upload()
	} */

	for i := 0; i < 34; i++ {
		project := models.CreateProject(
			"this is a test title for a project",
			"this is a test description for a project",
			"this is a test body text for a project",
			"https://github.com/Eiliv17",
			[]string{"html", "css", "go"},
		)
		project.Upload()
	}

}
