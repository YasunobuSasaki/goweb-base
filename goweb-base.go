package main

import (
 "github.com/gin-gonic/gin"
 "net/http"
)


func main() {
	r := gin.Default()	
	r.LoadHTMLGlob("templates/*")
	r.GET("/hello", func(c *gin.Context) {
		c.String(200, "hello")
	})

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "INDEX_PAGE",
		})
	})

	// 8080他で使ってたので5000
	r.Run(":5000")
}
