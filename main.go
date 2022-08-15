package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{"1", "Resah", "Payung Teduh", 50.500},
	{"2", "Membasuh", "Hindia", 35.999},
	{"3", "Senja di Ambang Pilu", "Danilla", 78.999},
	{"4", "Menuju Senja", "Payung Teduh", 45.788},
}

func main() {
	router := gin.Default()
	router.GET("/", sayHello)
	router.GET("/albums", getAlbums)
	router.GET("albums/:id", getAlbumsByID)
	router.DELETE("albums/:id", deleteAlbumsByID)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8081")
}

func sayHello(c *gin.Context) {
	c.JSONP(http.StatusOK, gin.H{
		"message": "Hello World",
	})
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumsByID(c *gin.Context) {
	id := c.Param("id")
	found := false

	for _, value := range albums {
		if value.ID == id {
			c.IndentedJSON(http.StatusOK, value)
			found = true
			return
		}
	}

	if found == false {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"message": "Albums not found",
		})
	}
}

func deleteAlbumsByID(c *gin.Context) {
	id := c.Param("id")
	found := false

	for key, value := range albums {
		if value.ID == id {
			albums = append(albums[:key], albums[key+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{
				"message": "Albums deleted !",
			})
			found = true
		}
	}

	if found == false {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"message": "Albums not found",
		})
	}
}
