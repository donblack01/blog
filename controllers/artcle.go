package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
)

func Artcle(this *gin.Context) {
	id := this.Param("id")
	html := "artcle.html"
	data := gin.H{"name": id}
	reg := regexp.MustCompile(`^[1-9][0-9]{0,5}$`)
	if !reg.MatchString(id) {
		html = "404.html"
		data = gin.H{}
	}
	this.HTML(http.StatusOK, html, data)
}
