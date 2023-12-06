package models

import (
	"book-organizer/forms"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Review struct {
	ID      primitive.ObjectID `json:"id" bson:"_id"`
	Point   int                `json:"point" bson:"point"`
	Comment string             `json:"comment" bson:"comment"`
}

type ReviewModel struct{}

func (r *ReviewModel) CreateReview(userID primitive.ObjectID, review forms.CreateReviewCommand) error {
	validate_err := ValidateStruct(review)
	if validate_err != nil {
		return validate_err
	}

	collection := dbConnect.Database(databaseName).Collection("books")
	bookID, _ := primitive.ObjectIDFromHex(review.BookID)
	filter := bson.D{{Key: "user_id", Value: userID}, {Key: "_id", Value: bookID}}

	var book Book
	err := collection.FindOne(nil, filter).Decode(&book)

	if err != nil {
		return err
	}

	newReview := Review{
		ID:      primitive.NewObjectID(),
		Point:   review.Point,
		Comment: review.Comment,
	}
	reviews := append(book.Reviews, newReview)
	err = updateBookReview(collection, reviews, filter)

	return err
}

func (r *ReviewModel) UpdateReview(userID primitive.ObjectID, review forms.UpdateReviewCommand) error {
	validate_err := ValidateStruct(review)
	if validate_err != nil {
		return validate_err
	}

	collection := dbConnect.Database(databaseName).Collection("books")
	bookID, _ := primitive.ObjectIDFromHex(review.BookID)
	filter := bson.D{{Key: "user_id", Value: userID}, {Key: "_id", Value: bookID}}

	var book Book
	err := collection.FindOne(nil, filter).Decode(&book)

	if err != nil {
		return err
	}

	reviews := book.Reviews
	reviewID, _ := primitive.ObjectIDFromHex(review.ID)
	index := getReviewIndexByID(reviews, reviewID)
	if index == -1 {
		return errors.New("Review not found")
	}

	updatedReview := Review{
		ID:      reviewID,
		Point:   review.Point,
		Comment: review.Comment,
	}
	reviews[index] = updatedReview
	err = updateBookReview(collection, reviews, filter)

	return err
}

func getReviewIndexByID(reviews []Review, id primitive.ObjectID) int {
	for i := 0; i < len(reviews); i++ {
		if reviews[i].ID == id {
			return i
		}
	}

	return -1
}

func updateBookReview(collection *mongo.Collection, reviews []Review, filter primitive.D) error {
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "reviews", Value: reviews}}}}
	_, err := collection.UpdateOne(nil, filter, update)

	return err
}
