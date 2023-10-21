package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/robel-yemane/rest-api/models"
	"github.com/robel-yemane/rest-api/storage"
)

// GetAlbums handles the GET /v1/albums route
func GetAlbums(c echo.Context, store storage.AlbumStorer) error {
	albums, err := store.GetAlbums()
	if err != nil {
		return c.JSON(http.StatusNotFound, "err:"+err.Error())
	}
	return c.JSON(http.StatusOK, albums)
}

// GetAlbumByID handles the GET /v1/albums/:id route
func GetAlbumByID(c echo.Context, store storage.AlbumStorer) error {
	id := c.Param("id")
	album, err := store.GetAlbumByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, "album not found")
	}
	return c.JSON(http.StatusOK, album)
}

// AddAlbum handles the POST /v1/albums route
func AddAlbum(c echo.Context, store storage.AlbumStorer) error {
	var newAlbum models.Album
	if err := c.Bind(&newAlbum); err != nil {
		return err
	}

	err := store.AddAlbum(newAlbum)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Failed to add album: "+err.Error())
	}

	return c.JSON(http.StatusCreated, newAlbum)
}
