package forms

type CreateBookCommand struct {
	Name         string   `json:"name" bson:"name" validate:"required"`
	Author       string   `json:"author" bson:"author" validate:"required"`
	PurchaseDate string   `json:"purchaseDate" bson:"purchaseDate" validate:"required"`
	StartReadAt  string   `json:"startReadAt" bson:"startReadAt"`
	FinishReadAt string   `json:"finishReadAt" bson:"finishReadAt"`
	Categories   []string `json:"categories" bson:"categories"`
}

type UpdateBookCommand struct {
	ID           string   `json:"id" bson:"id" validate:"required"`
	Name         string   `json:"name" bson:"name" validate:"required"`
	Author       string   `json:"author" bson:"author" validate:"required"`
	PurchaseDate string   `json:"purchase_date" bson:"purchase_date" validate:"required"`
	StartReadAt  string   `json:"start_read_at" bson:"start_read_at" validate:"required"`
	FinishReadAt string   `json:"finish_read_at" bson:"finish_read_at" validate:"required"`
	Categories   []string `json:"categories" bson:"categories"`
}
