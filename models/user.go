package models

import (
	"book-organizer/forms"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID         primitive.ObjectID `bson:"_id" json:"id"`
	Name       string             `json:"name" bson:"name"`
	Email      string             `json:"email" bson:"email"`
	Password   string             `json:"password" bson:"password"`
	IsVerified bool               `json:"is_verified" bson:"is_verified"`
}

type UserModel struct{}

func (u *UserModel) Signup(data forms.SignupUserCommand) error {
	collection := dbConnect.Database(databaseName).Collection("users")
	filter := bson.D{{Key: "email", Value: data.Email}}
	count, _ := collection.CountDocuments(nil, filter)

	if count > 0 {
		return errors.New("Email registered")
	}

	hashed_password, _ := HashPassword(data.Password)
	_, err := collection.InsertOne(nil, bson.M{
		"name":        data.Name,
		"email":       data.Email,
		"password":    hashed_password,
		"is_verified": false,
	})

	return err
}

func (u *UserModel) SignIn(data forms.SignInUserCommand) (string, error) {
	collection := dbConnect.Database(databaseName).Collection("users")
	filter := bson.D{{Key: "email", Value: data.Email}}
	var user User
	err := collection.FindOne(context.TODO(), filter).Decode(&user)

	if err != nil {
		return "", err
	}

	if !user.IsVerified {
		err = errors.New("Email not verify")

		return "", err
	}

	if !CheckPasswordHash(data.Password, user.Password) {
		err = errors.New("Email or password incorrect")

		return "", err
	}

	var session SessionModel
	token, err := session.CreateSession(user.ID)

	return token, err
}

func (u *UserModel) UpdatePassword(data forms.UpdateUserPasswordCommand) error {
	collection := dbConnect.Database(databaseName).Collection("users")
	id, _ := primitive.ObjectIDFromHex(data.ID)
	filter := bson.D{{Key: "_id", Value: id}}
	if !ValidOldPassword(collection, data, filter) {
		return errors.New("Password Incorrect")
	}

	new_password, _ := HashPassword(data.NewPassword)
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "password", Value: new_password}}}}

	_, err := collection.UpdateOne(nil, filter, update)

	return err
}

func (u *UserModel) FindOne(id primitive.ObjectID) (User, error) {
	collection := dbConnect.Database(databaseName).Collection("users")
	filter := bson.D{{Key: "_id", Value: id}}
	var user User
	err := collection.FindOne(nil, filter).Decode(&user)

	return user, err
}

func (u *UserModel) FindOneByEmail(email string) (User, error) {
	collection := dbConnect.Database(databaseName).Collection("users")
	filter := bson.D{{Key: "email", Value: email}}
	var user User
	err := collection.FindOne(nil, filter).Decode(&user)

	return user, err
}

func ValidOldPassword(collection *mongo.Collection, data forms.UpdateUserPasswordCommand, filter bson.D) bool {
	var user User
	err := collection.FindOne(nil, filter).Decode(&user)
	if err != nil {
		return false
	}

	return CheckPasswordHash(data.Password, user.Password)
}

func (u *UserModel) VerifyAccount(email string) error {
	collection := dbConnect.Database(databaseName).Collection("users")
	filter := bson.D{{Key: "email", Value: email}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "is_verified", Value: true},
		}},
	}

	_, err := collection.UpdateOne(nil, filter, update)

	return err
}
