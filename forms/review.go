package forms

type CreateReviewCommand struct {
	BookID  string `bson:"book_id" json:"book_id" validate:"required"`
	Point   int    `bson:"point" json:"point" validate:"required, gte=1, lte=5"`
	Comment string `bson:"comment" json:"comment" validate:"required, len=500"`
}

type UpdateReviewCommand struct {
	BookID  string `bson:"book_id" json:"book_id" validate:"required"`
	ID      string `bson:"id" json:"id" validate:"required"`
	Point   int    `bson:"point" json:"point" validate:"gte=1,lte=5"`
	Comment string `bson:"comment" json:"comment" validate:"required,max=500"`
}
