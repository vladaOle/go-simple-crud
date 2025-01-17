package repositories

import (
    "context"
    "encoding/json"
    "github.com/go-redis/redis/v8"
    "go-simple-crud/models"
    "strconv"
    "errors"
)

type RedisAlbumRepository struct {
    Client *redis.Client
    Ctx    context.Context
}

func (r *RedisAlbumRepository) Create(album models.Album) (models.Album, error) {
    albumKey := "album:" + strconv.Itoa(album.ID)
    albumJSON, err := json.Marshal(album)
    if err != nil {
        return models.Album{}, err
    }
    err = r.Client.Set(r.Ctx, albumKey, albumJSON, 0).Err()
    if err != nil {
        return models.Album{}, err
    }
    return album, nil
}

func (r *RedisAlbumRepository) GetList() ([]models.Album, error) {
    var albums []models.Album

    keys, err := r.Client.Keys(r.Ctx, "album:*").Result()
    if err != nil {
        return nil, err
    }

    for _, key := range keys {
        albumJSON, err := r.Client.Get(r.Ctx, key).Result()
        if err != nil {
            continue
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

    exists, err := r.Client.Exists(r.Ctx, albumKey).Result()
    if err != nil {
        return models.Album{}, err
    }
    
    if exists == 0 {
        return models.Album{}, errors.New("album not found")
    }

    albumJSON, err := json.Marshal(album)
    if err != nil {
        return models.Album{}, err
    }
    err = r.Client.Set(r.Ctx, albumKey, albumJSON, 0).Err()
    if err != nil {
        return models.Album{}, err
    }
    return album, nil
}

func (r *RedisAlbumRepository) Delete(id int) error {
    albumKey := "album:" + strconv.Itoa(id)

    exists, err := r.Client.Exists(r.Ctx, albumKey).Result()
    if err != nil {
        return err
    }

    if exists == 0 {
        return errors.New("album not found")
    }

    err = r.Client.Del(r.Ctx, albumKey).Err()
    return err
}