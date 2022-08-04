package dto

// LoginDTO is used when client post from /login url
type LoginDTO struct {
	Email    string `json:"email" from:"email" binding:"required,email"`
	Password string `json:"password" from:"password" binding:"required"`
}
