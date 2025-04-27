package services

import (
	"log"

	"github.com/adiva2311/product_api.git/helpers"
	"github.com/adiva2311/product_api.git/models"
	"github.com/adiva2311/product_api.git/repositories"
	"gorm.io/gorm"
)

type ProductService interface {
	Create(request models.Product) helpers.ProductResponse
	GetByUserId(product_id int, user_id int) helpers.ProductResponse
	Update(product_id int, user_id int, request models.Product) helpers.ProductResponse
	Delete(product_id int, user_id int)
}

type ProductServiceImpl struct {
	ProductRepo repositories.ProductRepository
}

// Create implements ProductService.
func (p *ProductServiceImpl) Create(request models.Product) helpers.ProductResponse {
	product := &models.Product{
		ProductName: request.ProductName,
		Total:       request.Total,
		Price:       request.Price,
		UserID:      request.UserID,
	}

	err := p.ProductRepo.Create(*product)
	if err != nil {
		log.Fatal("Failed to Add Product")
	}

	return helpers.ToProductResponse(*product)
}

// Delete implements ProductService.
func (p *ProductServiceImpl) Delete(product_id int, user_id int) {
	panic("unimplemented")
}

// GetByUserId implements ProductService.
func (p *ProductServiceImpl) GetByUserId(product_id int, user_id int) helpers.ProductResponse {
	panic("unimplemented")
}

// Update implements ProductService.
func (p *ProductServiceImpl) Update(product_id int, user_id int, request models.Product) helpers.ProductResponse {
	panic("unimplemented")
}

func NewProductService(db *gorm.DB) ProductService {
	return &ProductServiceImpl{
		ProductRepo: repositories.NewProductRepository(db),
	}
}
