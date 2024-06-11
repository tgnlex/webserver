package main

import "github.com/gin-gonic/gin"

func viewHandler(c *gin.Context) {
	c.HTML(ok, "index.tmpl", gin.H{
		"title": "Home",
	})
}
