package models

import (
	"book-organizer/forms"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	ID                 primitive.ObjectID `bson:"_id" json:"id"`
	Name               string             `json:"name" bson:"name"`
	Email              string             `json:"email" bson:"email"`
	Password           string             `json:"-" bson:"password"`
	IsVerified         bool               `json:"-" bson:"is_verified"`
	AvatarKey          string             `json:"avatar_key" bson:"avatar_key"`
	AvatarUrl          string             `json:"avatar_url" bson:"-"`
	FavoriteCategories []string           `json:"favorite_categories" bson:"favorite_categories"`
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

	newPassword, _ := HashPassword(data.NewPassword)
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "password", Value: newPassword}}}}

	session, session_err := dbConnect.StartSession()
	if session_err != nil {
		return session_err
	}
	defer session.EndSession(context.TODO())

	if transaction_err := session.StartTransaction(); transaction_err != nil {
		return transaction_err
	}

	var err error
	if err = mongo.WithSession(context.TODO(), session, func(sc mongo.SessionContext) error {
		_, err := collection.UpdateOne(nil, filter, update)
		sessionModel := new(SessionModel)
		err = sessionModel.DeleteAllUserSession(id)
		if err != nil {
			return err
		}

		if err = session.CommitTransaction(sc); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return err
}

func (u *UserModel) UpdateAvatar(data forms.UpdateAvatarCommand, id primitive.ObjectID) error {
	collection := dbConnect.Database(databaseName).Collection("users")
	filter := bson.D{{Key: "_id", Value: id}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "avatar_key", Value: data.Key}}}}
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

func (u *UserModel) AvatarKeys() []string {
	collection := dbConnect.Database(databaseName).Collection("users")
	var avatarKeys []string

	opts := options.Find().SetProjection(bson.D{{Key: "avatar_key", Value: 1}})
	filter := bson.D{{Key: "avatar_key", Value: bson.D{{Key: "$exists", Value: true}}}}
	cursor, err := collection.Find(context.TODO(), filter, opts)
	if err != nil {
		return avatarKeys
	}

	var users []User
	if err = cursor.All(context.TODO(), &users); err != nil {
		return avatarKeys
	}

	for _, user := range users {
		if user.AvatarKey != "" {
			avatarKeys = append(avatarKeys, user.AvatarKey)
		}
	}

	return avatarKeys
}

func (u *UserModel) UpdateUserFavoriteCategories(results []bson.M, userID primitive.ObjectID) error {
	collection := dbConnect.Database(databaseName).Collection("users")
	favoriteCategories := u.getMostFavoriteCategories(results, 5)

	filter := bson.D{{Key: "_id", Value: userID}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "favorite_categories", Value: favoriteCategories}}}}
	_, err := collection.UpdateOne(context.TODO(), filter, update)

	return err
}

func (u *UserModel) GetFavoriteCategories() ([]string, error) {
	collection := dbConnect.Database(databaseName).Collection("users")
	results, err := collection.Distinct(context.TODO(), "favorite_categories", bson.D{})
	var formatedResults []string
	for _, result := range results {
		formatedResult := result.(string)
		formatedResults = append(formatedResults, formatedResult)
	}

	return formatedResults, err
}

func (u *UserModel) getMostFavoriteCategories(results []bson.M, maxFavoriteCategories int) []string {
	favoriteCategories := []string{}

	for _, result := range results {
		if maxFavoriteCategories < 1 {
			break
		}

		value := result["categories"].(string)
		favoriteCategories = append(favoriteCategories, value)
		maxFavoriteCategories -= 1
	}

	return favoriteCategories
}
