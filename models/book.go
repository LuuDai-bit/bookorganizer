package models

import (
	"book-organizer/forms"
	"book-organizer/services"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BookType string

const (
	paperback BookType = "paperback"
	ebook     BookType = "ebook"
)

type Book struct {
	ID           primitive.ObjectID `bson:"_id" json:"id"`
	UserID       primitive.ObjectID `bson:"user_id" json:"user_id"`
	Name         string             `json:"name" bson:"name"`
	Author       string             `json:"author" bson:"author"`
	PurchaseDate primitive.DateTime `json:"purchase_date,omitempty" bson:"purchase_date,omitempty"`
	StartReadAt  primitive.DateTime `json:"start_read_at,omitempty" bson:"start_read_at,omitempty"`
	FinishReadAt primitive.DateTime `json:"finish_read_at,omitempty" bson:"finish_read_at,omitempty"`
	FileName     string             `json:"file_name" bson:"file_name"`
	Key          string             `json:"-" bson:"key,omitempty"`
	Type         BookType           `json:"type,omitempty" bson:"type,omitempty" default:"normal"`
	Categories   []string           `json:"categories" bson:"categories"`
	Reviews      []Review           `json:"reviews" bson:"reviews"`
}

type BookModel struct{}

func (b *BookModel) GetBooksByUser(userID primitive.ObjectID, page int, search string) ([]Book, error) {
	collection := dbConnect.Database(databaseName).Collection("books")

	filterStage := bson.D{{Key: "user_id", Value: userID}}
	if search != "" {
		filterStage = append(bson.D{{Key: "$text", Value: bson.D{{Key: "$search", Value: search}}}}, filterStage...)
	}

	filter := bson.D{{Key: "$match", Value: filterStage}}
	sort := bson.D{{Key: "$sort", Value: bson.D{{Key: "purchase_date", Value: -1}}}}
	skip := bson.D{{Key: "$skip", Value: (page - 1) * 20}}
	limit := bson.D{{Key: "$limit", Value: 20}}

	cursor, err := collection.Aggregate(nil, mongo.Pipeline{filter, sort, skip, limit})
	if err != nil {
		return nil, err
	}

	var books []Book
	if err = cursor.All(nil, &books); err != nil {
		return nil, err
	}

	return books, err
}

func (b *BookModel) GetTotalBookByUser(userID primitive.ObjectID, search string) (int64, error) {
	collection := dbConnect.Database(databaseName).Collection("books")
	filter := bson.D{{Key: "user_id", Value: userID}}
	total, err := collection.CountDocuments(context.TODO(), filter)

	return total, err
}

func (b *BookModel) CreateBook(userID primitive.ObjectID, book forms.CreateBookCommand) error {
	validate_err := ValidateStruct(book)
	if validate_err != nil {
		return validate_err
	}

	collection := dbConnect.Database(databaseName).Collection("books")

	PPurchaseDate, PStartReadAt, PFinishReadAt, err := ConvertBookDateField(book.PurchaseDate, book.StartReadAt, book.FinishReadAt)
	if err != nil {
		return err
	}

	_, err = collection.InsertOne(nil, bson.M{
		"user_id":        userID,
		"name":           book.Name,
		"author":         book.Author,
		"purchase_date":  PPurchaseDate,
		"start_read_at":  PStartReadAt,
		"finish_read_at": PFinishReadAt,
		"categories":     book.Categories,
		"file_name":      book.FileName,
		"key":            book.Key,
		"type":           book.Type,
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
	PPurchaseDate, PStartReadAt, PFinishReadAt, err := ConvertBookDateField(book.PurchaseDate, book.StartReadAt, book.FinishReadAt)
	if err != nil {
		return err
	}

	update := bson.D{{Key: "$set", Value: bson.D{
		{Key: "name", Value: book.Name},
		{Key: "author", Value: book.Author},
		{Key: "purchase_date", Value: PPurchaseDate},
		{Key: "start_read_at", Value: PStartReadAt},
		{Key: "finish_read_at", Value: PFinishReadAt},
		{Key: "categories", Value: book.Categories},
		{Key: "key", Value: book.Key},
		{Key: "file_name", Value: book.FileName},
		{Key: "type", Value: book.Type},
	}}}

	_, err = collection.UpdateOne(nil, filter, update)

	return err
}

func (b *BookModel) DownloadEbook(userID primitive.ObjectID, id primitive.ObjectID) (string, error) {
	collection := dbConnect.Database(databaseName).Collection("books")
	filter := bson.D{{Key: "user_id", Value: userID}, {Key: "_id", Value: id}}

	var book Book
	collection.FindOne(nil, filter).Decode(&book)

	s3Handler := new(services.S3Handler)
	url, err := s3Handler.GeneratePresignUrl(book.Key)

	return url, err
}

func ConvertBookDateField(purchaseDate string, startReadAt string, finishReadAt string) (*time.Time, *time.Time, *time.Time, error) {
	purchaseDateConveted, errPD := time.Parse(time.DateOnly, purchaseDate)
	startReadAtConverted, errSRA := time.Parse(time.DateTime, startReadAt)
	finishReadAtConverted, errFRA := time.Parse(time.DateTime, finishReadAt)

	var PPurchaseDate, PStartReadAt, PFinishReadAt *time.Time
	if errPD != nil {
		PPurchaseDate = nil
	} else {
		PPurchaseDate = &purchaseDateConveted
	}

	if errSRA != nil {
		PStartReadAt = nil
	} else {
		PStartReadAt = &startReadAtConverted
	}

	if errFRA != nil {
		PFinishReadAt = nil
	} else {
		PFinishReadAt = &finishReadAtConverted
	}

	var err error
	switch {
	case purchaseDate != "" && errPD != nil:
		err = errPD
	case startReadAt != "" && errSRA != nil:
		err = errSRA
	case finishReadAt != "" && errFRA != nil:
		err = errFRA
	}

	return PPurchaseDate, PStartReadAt, PFinishReadAt, err
}

func (u *BookModel) EbookKeys() ([]string, error) {
	collection := dbConnect.Database(databaseName).Collection("books")

	opts := options.Find().SetProjection(bson.D{{Key: "key", Value: 1}})
	filter := bson.D{{Key: "key", Value: bson.D{{Key: "$exists", Value: true}}}}
	cursor, err := collection.Find(context.TODO(), filter, opts)
	if err != nil {
		return nil, err
	}

	var books []Book
	if err = cursor.All(context.TODO(), &books); err != nil {
		return nil, err
	}

	var ebookKeys []string
	for _, book := range books {
		if book.Key != "" {
			ebookKeys = append(ebookKeys, book.Key)
		}
	}

	return ebookKeys, err
}
