package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Calendar struct {
	ID           primitive.ObjectID `bson:"_id" json:"id"`
	UserID       primitive.ObjectID `bson:"user_id" json:"user_id"`
	EventDate    primitive.DateTime `bson:"event_date" json:"event_date"`
	EventDetails []EventDetail      `bson:"event_details" json:"event_details"`
}

type EventDetail struct {
	Time             primitive.DateTime `bson:"time" json:"time"`
	ShortDescription string             `bson:"short_description" json:"short_description"`
	Detail           string             `bson:"detail" json:"detail"`
}

type CalendarModel struct{}

func (c *CalendarModel) GetCalendars(userID primitive.ObjectID, startDate time.Time, endDate time.Time) ([]Calendar, error) {
	collection := dbConnect.Database(databaseName).Collection("calendars")

	filter := bson.D{
		{Key: "user_id", Value: userID},
		{Key: "EventDate", Value: bson.D{
			{Key: "$gte", Value: startDate},
			{Key: "$lt", Value: endDate},
		}},
	}

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	var calendars []Calendar
	if err = cursor.All(context.TODO(), &calendars); err != nil {
		return nil, err
	}

	return calendars, nil
}
