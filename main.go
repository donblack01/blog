package main

import (
	"gin/controllers"
	"github.com/gin-gonic/gin"
	"os"
	"path/filepath"
)

func main() {
	blog := gin.Default()
	blog.Use(gin.Logger())
	gin.SetMode(gin.DebugMode)
	blog.Delims("{{", "}}")
	blog.LoadHTMLGlob(filepath.Join(os.Getenv("GOPATH"), "src/gin/views/*"))
	blog.GET("/", controllers.Index)
	blog.Run(":8080")
}
