package repositories

import (
	"product_api/models"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(product *models.Product) error
	GetByUserId(user_id int) ([]models.Product, error)
	Update(product_id int, user_id int, product *models.Product) error
	Delete(product_id int, user_id int) (int, error)
}

type ProductRepositoryImpl struct {
	Conn *gorm.DB
}

// Create implements ProductRepository.
func (p *ProductRepositoryImpl) Create(product *models.Product) error {
	return p.Conn.Create(product).Error
}

// Delete implements ProductRepository.
func (p *ProductRepositoryImpl) Delete(product_id int, user_id int) (int, error) {
	//return p.Conn.Delete(&models.Product{}, product_id, user_id).Error
	//return p.Conn.Where("id = ? AND user_id = ?", product_id, user_id).Delete(&models.Product{}).Error
	result := p.Conn.Where("id = ? AND user_id = ?", product_id, user_id).Delete(&models.Product{})
	rowsAffected := result.RowsAffected
	if result.Error != nil {
		return 0, result.Error
	}
	return int(rowsAffected), nil
}

// GetAll implements ProductRepository.
func (p *ProductRepositoryImpl) GetByUserId(user_id int) ([]models.Product, error) {
	var data []models.Product
	result := p.Conn.Where("user_id = ?", user_id).Find(&data)
	return data, result.Error
}

// Update implements ProductRepository.
func (p *ProductRepositoryImpl) Update(product_id int, user_id int, product *models.Product) error {
	return p.Conn.Where("id = ? AND user_id = ?", product_id, user_id).Updates(product).Error
}

func NewProductRepository(conn *gorm.DB) ProductRepository {
	return &ProductRepositoryImpl{Conn: conn}
}
