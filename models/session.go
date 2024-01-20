package models

import (
	"time"

	"github.com/google/uuid"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Session struct {
	ExpireTime primitive.DateTime `json:"-" bson:"expire_time"`
	Token      string             `json:"token" bson:"token"`
	UserID     primitive.ObjectID `json:"user_id" bson:"user_id"`
}

type SessionModel struct{}

func (s *SessionModel) CreateSession(userID primitive.ObjectID) (string, error) {
	collection := dbConnect.Database(databaseName).Collection("sessions")

	token := uuid.New().String()
	_, err := collection.InsertOne(nil, bson.M{
		"expire_time": time.Now().Add(24 * time.Hour),
		"token":       token,
		"user_id":     userID,
	})

	return token, err
}

func (s *SessionModel) Destroy(token string) (*mongo.DeleteResult, error) {
	collection := dbConnect.Database(databaseName).Collection("sessions")

	filter := bson.D{
		{Key: "token", Value: token},
		{Key: "expire_time", Value: bson.D{
			{Key: "$gt", Value: time.Now()},
		}},
	}
	result, err := collection.DeleteOne(nil, filter)

	return result, err
}

func (s *SessionModel) FindOne(token string) (Session, error) {
	collection := dbConnect.Database(databaseName).Collection("sessions")

	filter := bson.D{{Key: "token", Value: token}}
	var session Session
	err := collection.FindOne(nil, filter).Decode(&session)

	return session, err
}

func (s *SessionModel) DeleteAllUserSession(userID primitive.ObjectID) error {
	collection := dbConnect.Database(databaseName).Collection("sessions")
	filter := bson.D{{Key: "user_id", Value: userID}}
	_, err := collection.DeleteMany(nil, filter)

	return err
}
