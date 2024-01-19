package models

import (
	"book-organizer/services"
	"context"
	"slices"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (f *FileModel) RemoveUnusedFile() {
	collection := dbConnect.Database(databaseName).Collection("files")

	userModel := new(UserModel)
	userKeys := userModel.AvatarKeys()

	bookModel := new(BookModel)
	bookKeys := bookModel.EbookKeys()

	usedKeys := append(userKeys, bookKeys...)

	opts := options.Find().SetProjection(bson.D{{Key: "key", Value: 1}})
	cursor, err := collection.Find(context.TODO(), bson.D{}, opts)
	if err != nil {
		panic(err)
	}

	var files []File
	if err = cursor.All(context.TODO(), &files); err != nil {
		panic(err)
	}

	var keys []string
	for _, file := range files {
		keys = append(keys, file.Key)
	}

	size := len(keys) - len(usedKeys)
	unusedKeys := make([]string, size)
	for i := 0; i < len(keys); i++ {
		if !slices.Contains(usedKeys, keys[i]) {
			unusedKeys = append(unusedKeys, keys[i])
		}
	}

	s3Handler := new(services.S3Handler)
	successDeleteKeys := make([]string, size)
	for i := 0; i < len(unusedKeys); i++ {
		err = s3Handler.DeleteObject(unusedKeys[i])
		if err == nil {
			successDeleteKeys = append(successDeleteKeys, unusedKeys[i])
		}
	}
	filter := bson.D{{Key: "key", Value: bson.D{{Key: "$in", Value: successDeleteKeys}}}}
	collection.DeleteMany(context.TODO(), filter)
}
