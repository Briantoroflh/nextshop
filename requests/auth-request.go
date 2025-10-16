package requests

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterRequest struct {
	FullName     string `json:"full_name" binding:"required"`
	Email        string  `json:"email" binding:"required"`
	Phone        *string `json:"phone" binding:"required"`
	PasswordHash string  `json:"password" binding:"required"`
	Role         int     `json:"role" binding:"required"`
}