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
	GetByUserId(user_id int) []models.Product
	Update(product_id int, user_id int, request models.Product) helpers.ProductResponse
	Delete(product_id int, user_id int) error
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

	err := p.ProductRepo.Create(product)
	if err != nil {
		log.Fatal("Failed to Add Product")
	}

	return helpers.ToProductResponse(*product)
}

// Delete implements ProductService.
func (p *ProductServiceImpl) Delete(product_id int, user_id int) error {
	err := p.ProductRepo.Delete(product_id, user_id)
	if err != nil {
		log.Fatal("Failed to Delete Product")
	}
	return nil
}

// GetByUserId implements ProductService.
func (p *ProductServiceImpl) GetByUserId(user_id int) []models.Product {
	listProduct, err := p.ProductRepo.GetByUserId(user_id)
	if err != nil {
		log.Fatal("Failed to Get Product")
	}

	return listProduct
}

// Update implements ProductService.
func (p *ProductServiceImpl) Update(product_id int, user_id int, request models.Product) helpers.ProductResponse {
	product := &models.Product{
		ProductName: request.ProductName,
		Total:       request.Total,
		Price:       request.Price,
		UserID:      request.UserID,
	}

	err := p.ProductRepo.Update(product_id, user_id, product)
	if err != nil {
		log.Fatal("Failed to Update Product")
	}

	return helpers.ToProductResponse(*product)
}

func NewProductService(db *gorm.DB) ProductService {
	return &ProductServiceImpl{
		ProductRepo: repositories.NewProductRepository(db),
	}
}
