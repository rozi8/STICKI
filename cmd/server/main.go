package main

import (
	"net/http"

	"sticki-backend/internal/handler"
	"sticki-backend/internal/infrastructure/memory"
	"sticki-backend/internal/usecase"
)

func main() {

	// PRODUCT
	productRepo := memory.NewProductMemory()
	productUC := usecase.NewProductUsecase(productRepo)
	productHandler := handler.NewProductHandler(productUC)

	// API ROUTES
	http.HandleFunc("/api/products", productHandler.GetProducts)
	http.HandleFunc("/api/admin/sale", handler.CreateSale)

	// STATIC FRONTEND
	fs := http.FileServer(http.Dir("./web"))
	http.Handle("/", fs)

	println("STICKI server running on :8080")
	http.ListenAndServe(":8080", nil)
}
