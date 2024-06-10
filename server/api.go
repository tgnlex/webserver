package main

import (
	"github.com/gin-gonic/gin"
)

func fetchAlbums(c *gin.Context) {
	c.IndentedJSON(ok, albums)
}

func fetchAlbumByID(c *gin.Context) {
	id := c.Param("id")
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(ok, a)

		}
	}
}

func uploadAlbum(c *gin.Context) {
	var newAlbum album
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	albums = append(albums, newAlbum)
	c.IndentedJSON(httpStat, newAlbum)
}
