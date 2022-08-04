package entity

// User represents users table in database
type User struct {
	ID       string `gorm:"primary_key" json:"id"`
	Name     string `gorm:"type:varchar(255)" json:"name"`
	Email    string `gorm:"type:varchar(255)" json:"email"`
	Password string `gorm:"->;<-;not null" json:"-"`
	Token    string `gorm:"-" json:"token,omitempty"`
}
