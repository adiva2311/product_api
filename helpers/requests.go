package helpers

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ProductRequest struct {
	ProductName string  `json:"product_name"`
	Total       int     `json:"total"`
	Price       float32 `json:"price"`
	UserID      uint    `json:"user_id"`
}
