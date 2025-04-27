package helpers

import "github.com/adiva2311/product_api.git/models"

type ApiResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type UserResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

func ToRegisterResponse(user models.User) UserResponse {
	return UserResponse{
		Username: user.Username,
		Email:    user.Email,
		Role:     user.Role,
	}
}

type LoginResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Token    string `json:"token"`
}

func ToLoginResponse(user *models.User, token string) LoginResponse {
	return LoginResponse{
		Username: user.Username,
		Email:    user.Email,
		Role:     user.Role,
		Token:    token,
	}
}

type ProductResponse struct {
	ID          uint    `json:"id"`
	ProductName string  `json:"product_name"`
	Total       int     `json:"total"`
	Price       float32 `json:"price"`
	UserID      uint    `json:"user_id"`
}

func ToProductResponse(product models.Product) ProductResponse {
	return ProductResponse{
		ID:          product.ID,
		ProductName: product.ProductName,
		Total:       product.Total,
		Price:       product.Price,
		UserID:      product.UserID,
	}
}
