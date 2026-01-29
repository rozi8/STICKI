package handler

import (
	"encoding/json"
	"net/http"
	"sticki-backend/internal/domain"
)

var sales []domain.Sale

func CreateSale(w http.ResponseWriter, r *http.Request) {
	var s domain.Sale
	json.NewDecoder(r.Body).Decode(&s)
	sales = append(sales, s)
	w.WriteHeader(http.StatusCreated)
}
