package forms

type CreateBookCommand struct {
	Name         string   `json:"name" bson:"name"`
	Author       string   `json:"author" bson:"author"`
	PurchaseDate string   `json:"purchase_date" bson:"purchase_date"`
	StartReadAt  string   `json:"start_read_at" bson:"start_read_at"`
	FinishReadAt string   `json:"finish_read_at" bson:"finish_read_at"`
	Categories   []string `json:"categories" bson:"categories"`
}
