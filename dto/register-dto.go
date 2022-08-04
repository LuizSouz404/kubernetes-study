package dto

// RegisterDTO is used when client post from /register url
type RegisterDTO struct {
	Name     string `json:"name" from:"name" binding:"required"`
	Email    string `json:"email" from:"email" binding:"required,email"`
	Password string `json:"password" from:"password" binding:"required"`
}
