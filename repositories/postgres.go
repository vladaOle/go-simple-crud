package repositories

import (
    "database/sql"
    "go-simple-crud/models"
    "errors"
)

type PostgresAlbumRepository struct {
    DB *sql.DB
}

func (r *PostgresAlbumRepository) Create(album models.Album) (models.Album, error) {
    query := "INSERT INTO albums (title, artist, price) VALUES ($1, $2, $3) RETURNING id"
    err := r.DB.QueryRow(query, album.Title, album.Artist, album.Price).Scan(&album.ID)
    if err != nil {
        return models.Album{}, err
    }
    return album, nil
}

func (r *PostgresAlbumRepository) GetList() ([]models.Album, error) {
    rows, err := r.DB.Query("SELECT id, title, artist, price FROM albums")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var albums []models.Album
    for rows.Next() {
        var album models.Album
        if err := rows.Scan(&album.ID, &album.Title, &album.Artist, &album.Price); err != nil {
            return nil, err
        }
        albums = append(albums, album)
    }
    return albums, nil
}

func (r *PostgresAlbumRepository) Update(album models.Album) (models.Album, error) {
    query := "UPDATE albums SET title = $1, artist = $2, price = $3 WHERE id = $4"
    result, err := r.DB.Exec(query, album.Title, album.Artist, album.Price, album.ID)
    if err != nil {
        return models.Album{}, err
    }
    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return models.Album{}, err
    }
    if rowsAffected == 0 {
        return models.Album{}, errors.New("album not found")
    }
    return album, nil
}

func (r *PostgresAlbumRepository) Delete(id int) error {
    query := "DELETE FROM albums WHERE id = $1"
    result, err := r.DB.Exec(query, id)
    if err != nil {
        return err
    }
    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return err
    }
    if rowsAffected == 0 {
        return errors.New("album not found")
    }
    return nil
}