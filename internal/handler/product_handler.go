package handler

import (
	"encoding/json"
	"net/http"
	"sticki-backend/internal/usecase"
)

type ProductHandler struct {
	uc *usecase.ProductUsecase
}

func NewProductHandler(u *usecase.ProductUsecase) *ProductHandler {
	return &ProductHandler{uc: u}
}

func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(h.uc.GetProducts())
}
