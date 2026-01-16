package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CategoryBookModel struct{}
type CategoryBook struct {
	ID       string       `json:"_id" bson:"_id"`
	Category string       `json:"category" bson:"category"`
	Books    []GoogleBook `json:"books" bson:"books"`
}
type GoogleBook struct {
	ID          string   `json:"_id" bson:"_id"`
	Kind        string   `json:"kind" bson:"kind"`
	Title       string   `json:"title" bson:"title"`
	SubTitle    string   `json:"subtitle" bson:"subtitle"`
	Authors     []string `json:"authors" bson:"author"`
	Description string   `json:"description" bson:"description"`
	Categories  []string `json:"categories" bson:"categories"`
}

func (c *CategoryBookModel) GetSuggestedBooks(favoriteCategories []string) ([]GoogleBook, error) {
	if favoriteCategories == nil || len(favoriteCategories) == 0 {
		return nil, nil
	}

	collection := dbConnect.Database(databaseName).Collection("category_books")
	filter := bson.D{
		{Key: "category", Value: bson.D{
			{Key: "$in", Value: favoriteCategories},
		}},
	}

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	var categoryBooks []CategoryBook
	if err = cursor.All(context.TODO(), &categoryBooks); err != nil {
		return nil, err
	}

	var books []GoogleBook
	for _, categoryBook := range categoryBooks {
		books = append(books, categoryBook.Books...)
	}

	var ids []string
	for i := len(books) - 1; i >= 0; i-- {
		if c.contains(ids, books[i].ID) {
			books = append(books[:i], books[i+1:]...)
		}
	}

	return books, err
}

func (c *CategoryBookModel) Create(category string, books []GoogleBook) error {
	for _, book := range books {
		validate_err := ValidateStruct(book)
		if validate_err != nil {
			return validate_err
		}
	}

	session, session_err := dbConnect.StartSession()
	if session_err != nil {
		return session_err
	}
	defer session.EndSession(context.TODO())

	if transaction_err := session.StartTransaction(); transaction_err != nil {
		return transaction_err
	}

	if err := mongo.WithSession(context.TODO(), session, func(sc mongo.SessionContext) error {
		collection := dbConnect.Database(databaseName).Collection("category_books")
		c.DeleteByCategory(category)
		_, err := collection.InsertOne(nil, bson.M{
			"category": category,
			"books":    books,
		})

		return err
	}); err != nil {
		return err
	}

	return nil
}

func (c *CategoryBookModel) DeleteByCategory(category string) error {
	collection := dbConnect.Database(databaseName).Collection("category_books")
	filter := bson.D{{Key: "category", Value: category}}
	_, err := collection.DeleteMany(context.TODO(), filter)

	return err
}

func (c *CategoryBookModel) contains(ids []string, targetId string) bool {
	for _, id := range ids {
		if id == targetId {
			return true
		}
	}

	return false
}
