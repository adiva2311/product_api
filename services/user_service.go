package services

import (
	"errors"
	"log"

	"product_api/helpers"
	"product_api/models"
	"product_api/repositories"
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
	// Check Username if Exist
	user, err := u.UserRepo.CheckUsername(request.Username)
	if err != nil {
		return helpers.LoginResponse{}, errors.New("invalid username or password")
	}

	// Check Password
	if !helpers.CheckPasswordHash(request.Password, user.Password) {
		return helpers.LoginResponse{}, errors.New("invalid username or password")
	}

	// Generate JWT
	token, err := helpers.GenerateJWT(int(user.ID), user.Username, user.Role)
	if err != nil {
		return helpers.LoginResponse{}, err
	}
	return helpers.ToLoginResponse(user, token), nil
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
		Role:     request.Role,
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
