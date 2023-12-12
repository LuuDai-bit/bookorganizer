package forms

type SendVerifyCodeCommand struct {
	Email string `json:"email" bson:"email"`
}

type VerifyAccountCommand struct {
	Email      string `json:"email" bson:"email" validate:"required"`
	VerifyCode string `json:"verify_code" bson:"verify_code" validate:"required"`
}
