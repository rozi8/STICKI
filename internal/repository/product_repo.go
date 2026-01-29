package repository

import "sticki-backend/internal/domain"

type ProductRepository interface {
	FindAll() []domain.Product
	Create(domain.Product)
}
