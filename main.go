package main

import (
	"net/http"
	"os"

	"github.com/robel-yemane/rest-api/types"

	echo "github.com/labstack/echo/v4"
)

// album slice to seed record album data.
var albums = []types.Album{
	{ID: "1", Title: "Ygermena'lo", Artist: "Yohannes Tikabo", Price: 30},
	{ID: "2", Title: "Semai", Artist: "Abraham Afewerki", Price: 40},
	{ID: "3", Title: "Lilo", Artist: "Temesgen Yared", Price: 50},
}

// // getAlbums responds with the list of all albums as JSON
func getAlbums(c echo.Context) error {
	return c.JSON(http.StatusOK, albums)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c echo.Context) error {
	id := c.Param("id")

	// loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			return c.JSON(http.StatusOK, a)

		}
	}
	return c.JSON(http.StatusNotFound, "album not found")

}

// postAlbums adds an album from JSON received in the request body
func postAlbums(c echo.Context) error {
	var newAlbum types.Album
	//Call BindJSON to bind the received JSON to newAlbum
	if err := c.Bind(&newAlbum); err != nil {
		return err
	}
	albums = append(albums, newAlbum)
	return c.JSON(http.StatusCreated, newAlbum)
}

func main() {
	port := getPort("APP_PORT")

	e := echo.New()
	e.GET("/albums", getAlbums)
	//Associate the /albums/:id path with the getAlbumByID function.
	// the colon preceding an item in the path signifies that the item is a path parameter.
	e.GET("/album/:id", getAlbumByID)
	e.POST("albums", postAlbums)
	e.Logger.Fatal(e.Start(":" + port))
}

func getPort(envvar string) string {
	port := os.Getenv(envvar)
	if port == "" {
		port = "8080"
	}
	return port
}
