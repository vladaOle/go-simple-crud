package controllers

import (
    "github.com/gin-gonic/gin"
    "strconv"
    "net/http"
    "go-simple-crud/models"
    "go-simple-crud/services"
)

func CreateAlbum(s *services.AlbumService) gin.HandlerFunc {
    return func(c *gin.Context) {
        var album models.Album
        if err := c.ShouldBindJSON(&album); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        createdAlbum, err := s.Repo.Create(album)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusCreated, gin.H{"album": createdAlbum})
    }
}

func GetAlbums(s *services.AlbumService) gin.HandlerFunc {
    return func(c *gin.Context) {
        albums, err := s.Repo.GetList()
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, gin.H{"albums": albums})
    }
}

func UpdateAlbum(s *services.AlbumService) gin.HandlerFunc {
    return func(c *gin.Context) {
        idStr := c.Param("id")
        id, err := strconv.Atoi(idStr)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
            return
        }

        var album models.Album
        if err := c.ShouldBindJSON(&album); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        album.ID = id

        updatedAlbum, err := s.Repo.Update(album)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, gin.H{"album": updatedAlbum})
    }
}

func DeleteAlbum(s *services.AlbumService) gin.HandlerFunc {
    return func(c *gin.Context) {
        idStr := c.Param("id")
        id, err := strconv.Atoi(idStr)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
            return
        }

        if err := s.Repo.Delete(id); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.Status(http.StatusOK)
    }
}