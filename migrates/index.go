package migrates

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type IndexMigration struct{}

func (i *IndexMigration) Run() {
	fmt.Println("Preparing indexes")

	// Mongodb will not create index if that index already exist
	createBookNameIndex()

	fmt.Println("Done")
}

func createBookNameIndex() {
	collection := dbConnect.Database(databaseName).Collection("books")

	indexModel := mongo.IndexModel{
		Keys: bson.D{{Key: "name", Value: "text"}},
	}
	collection.Indexes().CreateOne(context.TODO(), indexModel)
}
