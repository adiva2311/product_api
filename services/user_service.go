package services

import (
	"log"

	"github.com/adiva2311/product_api.git/helpers"
	"github.com/adiva2311/product_api.git/models"
	"github.com/adiva2311/product_api.git/repositories"
	"gorm.io/gorm"
)

type UserService interface {
	Register(request models.User) (helpers.UserResponse, error)
	Login(request helpers.LoginRequest) (helpers.LoginResponse, error)
}

type UserServiceImpl struct {
	UserRepo repositories.UserRepository
}

// Login implements UserService.
func (u *UserServiceImpl) Login(request helpers.LoginRequest) (helpers.LoginResponse, error) {
	panic("unimplemented")
}

// Register implements UserService.
func (u *UserServiceImpl) Register(request models.User) (helpers.UserResponse, error) {
	//var response helpers.ApiResponse

	// Hash Password
	hashPassword, err := helpers.HashPassword(request.Password)
	if err != nil {
		log.Fatal(err)
	}

	user := &models.User{
		Username: request.Username,
		Email:    request.Email,
		Password: hashPassword,
	}

	err = u.UserRepo.Register(*user)
	if err != nil {
		log.Fatal("Failed to Register User")
	}

	return helpers.ToRegisterResponse(*user), nil
}

func NewUserService(db *gorm.DB) UserService {
	return &UserServiceImpl{
		UserRepo: repositories.NewUserRepository(db),
	}
}
