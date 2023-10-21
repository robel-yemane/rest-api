package main

import (
	"os"

	"github.com/robel-yemane/rest-api/internal/handlers"
	"github.com/robel-yemane/rest-api/internal/storage"

	echo "github.com/labstack/echo/v4"
)

func main() {
	port := getPort("APP_PORT")
	store := storage.NewMongoStore()

	e := echo.New()

	e.GET("/albums", func(c echo.Context) error { return handlers.GetAlbums(c, store) })
	e.GET("/albums/:id", func(c echo.Context) error { return handlers.GetAlbumByID(c, store) })
	e.POST("/albums", func(c echo.Context) error { return handlers.AddAlbum(c, store) })

	e.Logger.Fatal(e.Start(":" + port))
}

func getPort(envvar string) string {
	port := os.Getenv(envvar)
	if port == "" {
		port = "8080"
	}
	return port
}
