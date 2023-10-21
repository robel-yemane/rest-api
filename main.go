package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/robel-yemane/rest-api/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	echo "github.com/labstack/echo/v4"
)

// var album = types.Album{
// 	ID:     "1",
// 	Title:  "Ygermenalo",
// 	Artist: "Yohannes Tikabo",
// 	Price:  30,
// }

var collection *mongo.Collection
var client *mongo.Client

func init() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	var err error
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to and pinged MongoDB!")

}

// // getAlbums responds with the list of all albums as JSON
func getAlbums(c echo.Context) error {
	var albums []types.Album
	cursor, err := collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return c.JSON(http.StatusNotFound, "err: "+err.Error())
	}
	if err = cursor.All(context.Background(), &albums); err != nil {
		return c.JSON(http.StatusNotFound, "err: "+err.Error())
	}
	return c.JSON(http.StatusOK, albums)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c echo.Context) error {
	id := c.Param("id")
	var album types.Album
	err := collection.FindOne(context.TODO(), bson.D{{Key: "id", Value: id}}).Decode(&album)
	if err != nil {

		return c.JSON(http.StatusNotFound, "album not found")
	}
	return c.JSON(http.StatusOK, album)
}

// postAlbums adds an album from JSON received in the request body
func postAlbums(c echo.Context) error {
	var newAlbum types.Album
	//Call BindJSON to bind the received JSON to newAlbum
	if err := c.Bind(&newAlbum); err != nil {
		return err
	}
	insertResult, err := collection.InsertOne(context.Background(), newAlbum)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Inserted a single document: ", insertResult.InsertedID)
	return c.JSON(http.StatusCreated, newAlbum)
}

func main() {
	port := getPort("APP_PORT")
	collection = client.Database("albumstore").Collection("albums")

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
