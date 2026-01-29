package domain

type Order struct {
	ID        int     `json:"id"`
	ProductID int     `json:"product_id"`
	Qty       int     `json:"qty"`
	Total     float64 `json:"total"`
}
