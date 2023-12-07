package models

import (
	"book-organizer/forms"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Book struct {
	ID           primitive.ObjectID `bson:"_id" json:"id"`
	UserID       primitive.ObjectID `bson:"user_id" json:"user_id"`
	Name         string             `json:"name" bson:"name"`
	Author       string             `json:"author" bson:"author"`
	PurchaseDate primitive.DateTime `json:"purchase_date" bson:"purchase_date"`
	StartReadAt  primitive.DateTime `json:"start_read_at" bson:"start_read_at"`
	FinishReadAt primitive.DateTime `json:"finish_read_at" bson:"finish_read_at"`
	Categories   []string           `json:"categories" bson:"categories"`
	Reviews      []Review           `json:"reviews" bson:"reviews"`
}

type BookModel struct{}

func (b *BookModel) GetBooksByUser(userID primitive.ObjectID, page int) []Book {
	collection := dbConnect.Database(databaseName).Collection("books")
	filter := bson.D{{Key: "$match", Value: bson.D{{Key: "user_id", Value: userID}}}}
	sort := bson.D{{Key: "$sort", Value: bson.D{{Key: "purchase_date", Value: -1}}}}
	cursor, err := collection.Aggregate(nil, mongo.Pipeline{filter, sort})
	if err != nil {
		panic(err)
	}

	var books []Book
	if err = cursor.All(nil, &books); err != nil {
		panic(err)
	}

	return books
}

func (b *BookModel) CreateBook(userID primitive.ObjectID, book forms.CreateBookCommand) error {
	validate_err := ValidateStruct(book)
	if validate_err != nil {
		return validate_err
	}

	collection := dbConnect.Database(databaseName).Collection("books")

	purchaseDate, err := time.Parse(time.DateOnly, book.PurchaseDate)
	if err != nil {
		return err
	}

	_, err = collection.InsertOne(nil, bson.M{
		"user_id":        userID,
		"name":           book.Name,
		"author":         book.Author,
		"purchase_date":  purchaseDate,
		"start_read_at":  nil,
		"finish_read_at": nil,
		"categories":     book.Categories,
	})

	return err
}

func (b *BookModel) UpdateBook(userID primitive.ObjectID, book forms.UpdateBookCommand) error {
	validate_err := ValidateStruct(book)
	if validate_err != nil {
		return validate_err
	}

	collection := dbConnect.Database(databaseName).Collection("books")
	bookID, _ := primitive.ObjectIDFromHex(book.ID)
	filter := bson.D{{Key: "user_id", Value: userID}, {Key: "_id", Value: bookID}}
	purchaseDate, startReadAt, finishReadAt, err := ConvertBookDateField(book)
	if err != nil {
		return err
	}

	update := bson.D{{Key: "$set", Value: bson.D{
		{Key: "name", Value: book.Name},
		{Key: "author", Value: book.Author},
		{Key: "purchase_date", Value: purchaseDate},
		{Key: "start_read_at", Value: startReadAt},
		{Key: "finish_read_at", Value: finishReadAt},
		{Key: "categories", Value: book.Categories},
	}}}

	_, err = collection.UpdateOne(nil, filter, update)

	return err
}

func ConvertBookDateField(book forms.UpdateBookCommand) (time.Time, time.Time, time.Time, error) {
	purchaseDate, errPD := time.Parse(time.DateOnly, book.PurchaseDate)
	startReadAt, errSRA := time.Parse(time.DateTime, book.StartReadAt)
	finishReadAt, errFRA := time.Parse(time.DateTime, book.FinishReadAt)

	var err error
	switch {
	case book.PurchaseDate != "" && errPD != nil:
		err = errPD
	case book.StartReadAt != "" && errSRA != nil:
		err = errSRA
	case book.FinishReadAt != "" && errFRA != nil:
		err = errFRA
	}

	return purchaseDate, startReadAt, finishReadAt, err
}
