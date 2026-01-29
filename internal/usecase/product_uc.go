package usecase

import (
	"sticki-backend/internal/domain"
	"sticki-backend/internal/repository"
)

type ProductUsecase struct {
	repo repository.ProductRepository
}

func NewProductUsecase(r repository.ProductRepository) *ProductUsecase {
	return &ProductUsecase{repo: r}
}

func (u *ProductUsecase) GetProducts() []domain.Product {
	return u.repo.FindAll()
}
