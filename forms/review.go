package forms

type CreateReviewCommand struct {
	BookID  string  `bson:"bookId" json:"bookId" validate:"required"`
	Point   float32 `bson:"point" json:"point" validate:"required"`
	Comment string  `bson:"comment" json:"comment"`
}

type UpdateReviewCommand struct {
	BookID  string  `bson:"bookId" json:"bookId" validate:"required"`
	ID      string  `bson:"id" json:"id" validate:"required"`
	Point   float32 `bson:"point" json:"point" validate:"required"`
	Comment string  `bson:"comment" json:"comment"`
}
