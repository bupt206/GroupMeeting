package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.Static("../static", "./static")
	r.LoadHTMLFiles("./templates/summary.html")
	r.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "summary.html", nil)
	})
	r.Run(":8080")
}
