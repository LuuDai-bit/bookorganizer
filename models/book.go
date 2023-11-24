package models

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Review struct {
	Point   int    `json:"point" bson:"point"`
	Comment string `json:"comment" bson:"comment"`
}

type Book struct {
	ID           primitive.ObjectID `bson:"_id" json:"id"`
	UserId       primitive.ObjectID `bson:"user_id" json:"user_id"`
	Name         string             `json:"name" bson:"name"`
	Author       string             `json:"author" bson:"author"`
	PurchaseDate primitive.DateTime `json:"purchase_date" bson:"purchase_date"`
	StartReadAt  primitive.DateTime `json:"start_read_at" bson:"start_read_at"`
	FinishReadAt primitive.DateTime `json:"finish_read_at" bson:"finish_read_at"`
	Categories   []string
	Reviews      []Review
}

type BookModel struct{}

func (b *BookModel) GetBooksByUser(user_id primitive.ObjectID, page int) []Book {
	collection := dbConnect.Database(databaseName).Collection("books")
	filter := bson.D{{"user_id", user_id}}
	cursor, err := collection.Find(nil, filter)
	if err != nil {
		panic(err)
	}

	var books []Book
	if err = cursor.All(nil, &books); err != nil {
		panic(err)
	}

	return books
}
