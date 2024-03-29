package forms

type CreateBookCommand struct {
	Name         string   `json:"name" bson:"name" validate:"required"`
	Author       string   `json:"author" bson:"author" validate:"required"`
	PurchaseDate string   `json:"purchaseDate" bson:"purchaseDate" validate:"required"`
	StartReadAt  string   `json:"startReadAt" bson:"startReadAt"`
	FinishReadAt string   `json:"finishReadAt" bson:"finishReadAt"`
	Key          string   `json:"key" bson:"key"`
	FileName     string   `json:"fileName" bson:"fileName"`
	Type         string   `json:"type" bson:"type"`
	Categories   []string `json:"categories" bson:"categories"`
}

type UpdateBookCommand struct {
	ID           string   `json:"id" bson:"id" validate:"required"`
	Name         string   `json:"name" bson:"name" validate:"required"`
	Author       string   `json:"author" bson:"author" validate:"required"`
	PurchaseDate string   `json:"purchaseDate" bson:"purchaseDate" validate:"required"`
	StartReadAt  string   `json:"startReadAt" bson:"startReadAt"`
	FinishReadAt string   `json:"finishReadAt" bson:"finishReadAt"`
	Key          string   `json:"key" bson:"key"`
	FileName     string   `json:"fileName" bson:"fileName"`
	Type         string   `json:"type" bson:"type"`
	Categories   []string `json:"categories" bson:"categories"`
}
