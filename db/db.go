package db

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo/readpref"
)

var server = os.Getenv("DB_URL")

func Connect() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	// Use env for host and port
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI(server))

	return client
}

// func UseCollection(database string, collection string, client *mongo.Client) *mongo.Collection {
//   return client.Database(database).Collection(collection)
// }
