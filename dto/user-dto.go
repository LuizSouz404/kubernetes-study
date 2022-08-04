package dto

// UserCreateDTO is used when register a user
type UserCreateDTO struct {
	Name     string `json:"name" from:"name" binding:"required"`
	Email    string `json:"email" from:"email" binding:"required,email"`
	Password string `json:"password,omitempty" from:"password,omitempty"`
}

// UserUpdateDTO is used by client when PUT update profile
type UserUpdateDTO struct {
	Id       string `json:"id" from:"id"`
	Name     string `json:"name" from:"name" binding:"required"`
	Email    string `json:"email" from:"email" binding:"required" validate:"email"`
	Password string `json:"password,omitempty" from:"password,omitempty" validate:"min:6"`
}
