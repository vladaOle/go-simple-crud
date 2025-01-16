package services

import "go-simple-crud/models"

type AlbumRepository interface {
    Create(album models.Album) (models.Album, error)
    GetList() ([]models.Album, error)
    Update(album models.Album) (models.Album, error)
    Delete(id int) error
}

type AlbumService struct {
    Repo AlbumRepository
}

func NewAlbumService(Repo AlbumRepository) *AlbumService {
    return &AlbumService{Repo: Repo}
}