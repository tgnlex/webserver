package main

import "github.com/gin-gonic/gin"


func server() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/", viewHandler)
	router.GET("/ping", ping)
	router.GET("/albums", fetchAlbums)
	router.GET("/albums/:id", fetchAlbumByID)
	router.POST("/albums", uploadAlbum)
	router.Run("localhost:4000")
}
