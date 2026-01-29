package domain

type Sale struct {
	ID      int     `json:"id"`
	Date    string  `json:"date"`
	Income  float64 `json:"income"`
	Expense float64 `json:"expense"`
}
