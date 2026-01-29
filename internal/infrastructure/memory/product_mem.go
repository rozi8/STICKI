package memory

import "sticki-backend/internal/domain"

type ProductMemory struct {
	data []domain.Product
}

func NewProductMemory() *ProductMemory {
	return &ProductMemory{
		data: []domain.Product{
			{ID: 1, Name: "STICKI Original", Price: 10000},
			{ID: 2, Name: "STICKI Pedas", Price: 12000},
		},
	}
}

func (p *ProductMemory) FindAll() []domain.Product {
	return p.data
}

func (p *ProductMemory) Create(prod domain.Product) {
	p.data = append(p.data, prod)
}
