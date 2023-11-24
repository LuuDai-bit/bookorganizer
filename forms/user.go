package forms

// SignupUserCommand defines user form struct
type SignupUserCommand struct {
	// binding:"required" ensures that the field is provided
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignInUserCommand struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LogOutUserCommand struct {
	Token string `json:"token" binding:"required"`
}

type UpdateUserPasswordCommand struct {
	ID          string `bson:"_id" json:"id"`
	Password    string `json:"password" bson:"password"`
	NewPassword string `json:"new_password" bson:"new_password"`
}
