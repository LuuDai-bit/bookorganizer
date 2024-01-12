package models

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Meta struct {
	// Implement later
}
type File struct {
	ID   primitive.ObjectID `json:"id" bson:"_id"`
	Key  string             `json:"key" bson:"key"`
	Meta Meta               `json:"meta" bson:"meta"`
}
type FileModel struct{}

func (f *FileModel) CreateFile(key string) error {
	collection := dbConnect.Database(databaseName).Collection("files")

	_, err := collection.InsertOne(nil, bson.M{
		"key": key,
	})

	return err
}
