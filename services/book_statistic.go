package services

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type BookStatistic struct{}

func (b *BookStatistic) NumberOfBookRead(year string) []bson.M {
	collection := prepareCollection("books")
	year_format := bson.D{
		{Key: "$dateToString", Value: bson.D{
			{Key: "date", Value: "$finish_read_at"},
			{Key: "format", Value: "%Y"},
		}},
	}
	attr := bson.D{{Key: "year", Value: year_format}}
	project := bson.D{{Key: "$project", Value: attr}}
	filter := bson.D{{Key: "$match", Value: bson.D{{Key: "year", Value: year}}}}
	count := bson.D{
		{Key: "$count", Value: "book_count"},
	}
	cursor, err := collection.Aggregate(nil, mongo.Pipeline{project, filter, count})
	if err != nil {
		panic(err)
	}

	var results []bson.M
	if err = cursor.All(nil, &results); err != nil {
		panic(err)
	}

	return results
}
