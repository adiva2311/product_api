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
}

func ToRegisterResponse(user models.User) UserResponse {
	return UserResponse{
		Username: user.Username,
		Email:    user.Email,
	}
}

type LoginResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}
