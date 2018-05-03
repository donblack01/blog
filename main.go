package main

import (
	"blog/controllers"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	blog := gin.Default()
	blog.Use(gin.Logger())
	gin.SetMode(gin.DebugMode)
	blog.Delims("{{", "}}")
	blog.LoadHTMLGlob(filepath.Join(os.Getenv("GOPATH"), "src/blog/views/*"))
	blog.Static("/static", filepath.Join(os.Getenv("GOPATH"), "src/blog/static/"))
	blog.GET("/", controllers.Index)
	blog.Run(":8081")
}
