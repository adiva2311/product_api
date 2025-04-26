package repositories

import (
	"github.com/adiva2311/product_api.git/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	CheckUsername(username string) (*models.User, error)
	Register(request models.User) error
}

type dbUser struct {
	Conn *gorm.DB
}

// CheckUsername implements UserRepository.
func (d *dbUser) CheckUsername(username string) (*models.User, error) {
	var user models.User
	err := d.Conn.Where("username = ?", username).First(&user).Error
	return &user, err
}

// Register implements UserRepository.
func (d *dbUser) Register(request models.User) error {
	return d.Conn.Create(&request).Error
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &dbUser{
		Conn: db,
	}
}
