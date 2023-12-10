package forms

type SendVerifyCodeCommand struct {
	Email string `json:"email" bson:"email"`
}
