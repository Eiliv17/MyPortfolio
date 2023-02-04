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

	/* for i := 0; i < 34; i++ {
		project := models.CreateProject(
			"this is a test title for a project",
			"this is a test description for a project",
			"this is a test body text for a project",
			"https://github.com/Eiliv17",
			[]string{"html", "css", "go"},
		)
		project.Upload()
	} */

	text := `
	<h2>Hello this is a test</h2>
    <p>Lorem ipsum, dolor sit amet consectetur adipisicing elit. Iure itaque voluptate optio excepturi asperiores provident, exercitationem nobis molestias nesciunt ullam vel commodi qui sapiente labore rem repellendus cumque inventore, assumenda eos perferendis. Quaerat cupiditate neque magni fugit consequuntur! Neque, architecto?</p>
    <p>Lorem ipsum, dolor sit amet consectetur adipisicing elit. Iure itaque voluptate optio excepturi asperiores provident, exercitationem nobis molestias nesciunt ullam vel commodi qui sapiente labore rem repellendus cumque inventore, assumenda eos perferendis. Quaerat cupiditate neque magni fugit consequuntur! Neque, architecto?</p>

    <h2>Hello this is a test</h2>
    <p>Lorem ipsum, dolor sit amet consectetur adipisicing elit. Iure itaque voluptate optio excepturi asperiores provident, exercitationem nobis molestias nesciunt ullam vel commodi qui sapiente labore rem repellendus cumque inventore, assumenda eos perferendis. Quaerat cupiditate neque magni fugit consequuntur! Neque, architecto?</p>
	`

	post := models.CreatePost("this is a better test for the dimensions of the post title", "a simple description", text, []string{"test", "test"})

	post.Upload()

}
