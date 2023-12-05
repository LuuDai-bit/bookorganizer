package models

import (
	"book-organizer/forms"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	filter := bson.D{{"user_id", userID}}
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
	filter := bson.D{{"user_id", userID}, {"_id", bookID}}
	purchaseDate, startReadAt, finishReadAt, err := ConvertBookDateField(book)
	if err != nil {
		return err
	}

	update := bson.D{{"$set", bson.D{
		{"name", book.Name},
		{"author", book.Author},
		{"purchase_date", purchaseDate},
		{"start_read_at", startReadAt},
		{"finish_read_at", finishReadAt},
		{"categories", book.Categories},
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
