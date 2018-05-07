package main

import (
	"blog/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	blog := gin.Default()
	blog.Use(gin.Logger())
	gin.SetMode(gin.DebugMode)
	blog.Delims("{{", "}}")
	blog.LoadHTMLGlob("./views/*")
	blog.Static("/static", "./static/")
	blog.NoRoute(func(this *gin.Context) {
		this.HTML(http.StatusNotFound, "404.html", gin.H{})
	})
	blog.GET("/", controllers.Index)
	blog.GET("/artcle/:id", controllers.Artcle)
	blog.Run(":8081")
}
