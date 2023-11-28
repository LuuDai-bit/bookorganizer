package models

import (
	"time"

	"github.com/google/uuid"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Session struct {
	ExpireTime primitive.DateTime `bson:"expire_time"`
	Token      string             `json:"token" bson:"token"`
	UserId     primitive.ObjectID `json:"user_id" bson:"user_id"`
}

type SessionModel struct{}

func (s *SessionModel) CreateSession(user_id primitive.ObjectID) (string, error) {
	collection := dbConnect.Database(databaseName).Collection("sessions")

	token := uuid.New().String()
	_, err := collection.InsertOne(nil, bson.M{
		"expire_time": time.Now().Add(24 * time.Hour),
		"token":       token,
		"user_id":     user_id,
	})

	return token, err
}

func (s *SessionModel) Destroy(token string) (*mongo.DeleteResult, error) {
	collection := dbConnect.Database(databaseName).Collection("sessions")

	filter := bson.D{{"token", token}}
	result, err := collection.DeleteOne(nil, filter)

	return result, err
}

func (s *SessionModel) FindOne(token string) (Session, error) {
	collection := dbConnect.Database(databaseName).Collection("sessions")

	filter := bson.D{{"token", token}}
	var session Session
	err := collection.FindOne(nil, filter).Decode(&session)

	return session, err
}
