package main

import (
	"net/http"
	"os"
	"portfolio-server/controllers"
	"portfolio-server/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.LoadDatabase()
}

func main() {

	// enabling different modes
	gin.DisableConsoleColor()
	// gin.SetMode(gin.ReleaseMode)

	// creating the gin engine
	r := gin.New()
	r.SetTrustedProxies(nil)

	/*
		// creating the loggin file.
		f, _ := os.Create("gin.log")

		logformat := gin.LogFormatter(func(param gin.LogFormatterParams) string {
			// your custom format
			return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
				param.ClientIP,
				param.TimeStamp.Format(time.RFC1123),
				param.Method,
				param.Path,
				param.Request.Proto,
				param.StatusCode,
				param.Latency,
				param.Request.UserAgent(),
				param.ErrorMessage,
			)
		})

		logconfig := gin.LoggerConfig{
			Formatter: logformat,
			Output:    io.MultiWriter(f),
		}

		r.Use(gin.LoggerWithConfig(logconfig))
	*/

	// use the recovery middleware
	r.Use(gin.Recovery())

	r.LoadHTMLGlob("views/*")
	r.Static("/public", "./public")

	googleAnalyticsID := os.Getenv("GOOGLE_ANALYTICS")

	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"ganalyticsid": googleAnalyticsID,
		})
	})

	r.StaticFile("favicon.ico", "./publicfiles/favicon.ico")
	r.StaticFile("cv-en.pdf", "./publicfiles/cv-en.pdf")
	r.StaticFile("cv-it.pdf", "./publicfiles/cv-it.pdf")

	api := r.Group("/api")

	api.GET("/posts", controllers.RetrievePosts)
	api.GET("/posts/:id", controllers.RetrievePost)

	api.GET("/projects", controllers.RetrieveProjects)
	api.GET("/projects/:id", controllers.RetrieveProject)

	api.POST("/contact", controllers.PostContact)
	api.OPTIONS("/contact", controllers.OptionsContact)

	//r.Run() // listen and serve on 0.0.0.0:8080

	port := os.Getenv("PORT")
	r.RunTLS(":"+port, "servercert.pem", "serverkey.key")

}
