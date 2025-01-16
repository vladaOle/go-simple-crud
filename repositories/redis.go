package repositories

import (
    "context"
    "encoding/json"
    "github.com/go-redis/redis/v8"
    "go-simple-crud/models"
    "strconv"
)

type RedisAlbumRepository struct {
    Client *redis.Client
    ctx    context.Context
}

func (r *RedisAlbumRepository) Create(album models.Album) (models.Album, error) {
    albumKey := "album:" + strconv.Itoa(album.ID)
    albumJSON, err := json.Marshal(album)
    if err != nil {
        return models.Album{}, err
    }
    err = r.Client.Set(r.ctx, albumKey, albumJSON, 0).Err()
    if err != nil {
        return models.Album{}, err
    }
    return album, nil
}

func (r *RedisAlbumRepository) GetList() ([]models.Album, error) {
    var albums []models.Album
    keys, err := r.Client.Keys(r.ctx, "album:*").Result()
    if err != nil {
        return nil, err
    }

    for _, key := range keys {
        albumJSON, err := r.Client.Get(r.ctx, key).Result()
        if err != nil {
            continue // Skip if there's an error
        }
        var album models.Album
        err = json.Unmarshal([]byte(albumJSON), &album)
        if err == nil {
            albums = append(albums, album)
        }
    }
    return albums, nil
}

func (r *RedisAlbumRepository) Update(album models.Album) (models.Album, error) {
    albumKey := "album:" + strconv.Itoa(album.ID)
    albumJSON, err := json.Marshal(album)
    if err != nil {
        return models.Album{}, err
    }
    err = r.Client.Set(r.ctx, albumKey, albumJSON, 0).Err()
    if err != nil {
        return models.Album{}, err
    }
    return album, nil
}

func (r *RedisAlbumRepository) Delete(id int) error {
    albumKey := "album:" + strconv.Itoa(id)
    err := r.Client.Del(r.ctx, albumKey).Err()
    return err
}