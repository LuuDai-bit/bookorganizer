package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CategoryBook struct{}
type GoogleBook struct {
	ID          string   `json:"_id" bson:"_id"`
	Kind        string   `json:"kind" bson:"kind"`
	Title       string   `json:"title" bson:"title"`
	SubTitle    string   `json:"subtitle" bson:"subtitle"`
	Authors     []string `json:"authors" bson:"author"`
	Description string   `json:"description" bson:"description"`
	Categories  []string `json:"categories" bson:"categories"`
}

func (c *CategoryBook) Create(category string, books []GoogleBook) error {
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

func (c *CategoryBook) DeleteByCategory(category string) error {
	collection := dbConnect.Database(databaseName).Collection("category_books")
	filter := bson.D{{Key: "category", Value: category}}
	_, err := collection.DeleteMany(context.TODO(), filter)

	return err
}
