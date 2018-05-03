package controllers

import (
	"blog/framework"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type UserMetaResult struct {
	Umeta_id   string
	User_id    string
	Meta_key   string
	Meta_value string
}

func Index(this *gin.Context) {
	var results []UserMetaResult
	rows, err := framework.Instance().Query("SELECT * FROM opus_usermeta")
	if err != nil {
		// proper error handling instead of panic in your app
	}
	for rows.Next() {
		var result UserMetaResult
		if err := rows.Scan(&result.Umeta_id, &result.User_id, &result.Meta_key, &result.Meta_value); err != nil {
			log.Println(err)
			return
		}
		results = append(results, result)
	}

	this.HTML(http.StatusOK, "index.html", gin.H{"name": "sdfasdfsad", "dat": results})

}
