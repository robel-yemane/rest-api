package storage

import (
	"context"
	"log"

	"github.com/robel-yemane/rest-api/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoStore struct {
	collection *mongo.Collection
}

func NewMongoStore() *MongoStore {
	//Initialise MongoDB client and collection here
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	collection := client.Database("albumstore").Collection("albums")
	return &MongoStore{collection: collection}
}

// GetAlbums returns all albums from the database
func (m *MongoStore) GetAlbums() ([]models.Album, error) {
	var albums []models.Album
	cursor, err := m.collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.TODO(), &albums); err != nil {
		return nil, err
	}
	return albums, nil
}

// GetAlbumByID returns a single album from the database
func (m *MongoStore) GetAlbumByID(id string) (models.Album, error) {
	var album models.Album
	err := m.collection.FindOne(context.TODO(), bson.M{"id": id}).Decode(&album)
	if err != nil {
		return models.Album{}, err
	}
	return album, nil
}

func (m *MongoStore) AddAlbum(album models.Album) error {
	_, err := m.collection.InsertOne(context.Background(), album)
	return err
}
