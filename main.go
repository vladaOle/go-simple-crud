package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

type album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}

var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, albums)
}

func addAlbum(c *gin.Context) {
    var newAlbum album

    if err := c.BindJSON(&newAlbum); err != nil {
        return
    }

    albums = append(albums, newAlbum)
    c.IndentedJSON(http.StatusCreated, newAlbum)
}

func updateAlbum(c *gin.Context) {
    id := c.Param("id")

	var updatedAlbum album

    if err := c.BindJSON(&updatedAlbum); err != nil {
        return
    }

    updatedTitle := updatedAlbum.Title
    updatedArtist := updatedAlbum.Artist
    updatedPrice := updatedAlbum.Price

	for i := range albums {
		if albums[i].ID == id {
			albums[i].Title = updatedTitle
			albums[i].Artist = updatedArtist
			albums[i].Price = updatedPrice
			c.JSON(http.StatusOK, gin.H{"success": "album updated"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "album not found"})	
}

func removeAlbum(c *gin.Context) {
    id := c.Param("id")
	
	for i := range albums {
		if albums[i].ID == id {
			albums = append(albums[:i], albums[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"success": "album deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "album not found"})	
}

func main() {
	r := gin.Default()
    r.GET("/albums", getAlbums)
    r.POST("/albums", addAlbum)
	r.PUT("/albums/:id", updateAlbum)
	r.DELETE("/albums/:id", removeAlbum)
	r.Run("localhost:8080")
}