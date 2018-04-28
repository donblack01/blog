package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(this *gin.Context) {

	this.HTML(http.StatusOK, "index.html", gin.H{"name": "ssss"})

}
