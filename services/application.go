package services

import (
	"book-organizer/db"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
)

func prepareCollection(collection string) *mongo.Collection {
	var dbConnect = db.Connect()
	var databaseName = os.Getenv("DATABASE_NAME")

	return dbConnect.Database(databaseName).Collection(collection)
}
