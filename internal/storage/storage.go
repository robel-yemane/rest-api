package storage

import "github.com/robel-yemane/rest-api/internal/models"

type AlbumStorer interface {
	GetAlbums() ([]models.Album, error)
	GetAlbumByID(id string) (models.Album, error)
	AddAlbum(album models.Album) error
}
