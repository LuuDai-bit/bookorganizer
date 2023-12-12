package models

import (
	"book-organizer/mails"
	"errors"
	"io"

	"crypto/rand"

	"go.mongodb.org/mongo-driver/bson"
)

const MaxLengthCode = 6
const otpChars = "1234567890"

type Verify struct {
	Email      string `json:"email" bson:"email"`
	VerifyCode bool   `json:"verify_code" bson:"verify_code"`
}

type VerifyModel struct{}

func (v *VerifyModel) SendVerifyCode(email string) error {
	user, err := new(UserModel).FindOneByEmail(email)

	if err != nil {
		return err
	}

	if user.IsVerified {
		err = errors.New("User email already verified")

		return err
	}

	v.deleteAllPreviousVerifyCode(email)
	verifyCode := v.generateVerifyCode(email)

	verifyMail := new(mails.VerifyMail)
	go verifyMail.SendVerifyCode(email, verifyCode)

	return err
}

func (v *VerifyModel) ValidateVerifyCode(email string, verifyCode string) error {
	collection := dbConnect.Database(databaseName).Collection("verify_code")
	filter := bson.D{
		{Key: "email", Value: email},
		{Key: "verify_code", Value: verifyCode},
	}
	var verifyRecord Verify
	err := collection.FindOne(nil, filter).Decode(&verifyRecord)

	if err != nil {
		return err
	}

	userModel := new(UserModel)
	userModel.VerifyAccount(email)
	v.deleteAllPreviousVerifyCode(email)

	return err
}

func (v *VerifyModel) deleteAllPreviousVerifyCode(email string) {
	collection := dbConnect.Database(databaseName).Collection("verify_code")
	filter := bson.D{{Key: "email", Value: email}}
	_, err := collection.DeleteMany(nil, filter)

	if err != nil {
		panic(err)
	}
}

func (v *VerifyModel) generateVerifyCode(email string) string {
	verifyCode := make([]byte, MaxLengthCode)
	_, err := io.ReadAtLeast(rand.Reader, verifyCode, MaxLengthCode)
	for i := 0; i < len(verifyCode); i++ {
		verifyCode[i] = otpChars[int(verifyCode[i])%MaxLengthCode]
	}

	verifyRecord := bson.M{
		"email":       email,
		"verify_code": string(verifyCode),
	}

	collection := dbConnect.Database(databaseName).Collection("verify_code")
	_, err = collection.InsertOne(nil, verifyRecord)

	if err != nil {
		panic(err)
	}

	return string(verifyCode)
}
