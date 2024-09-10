package Model

import "github.com/google/uuid"

type Auth struct {
	Id       uuid.UUID `gorm:"type:uuid;default:null" json:"id"`
	UserName string    `json:"userName"`
	Email    string    `json:"email"`
	Password string    `json:"-"`
	Role     string    ` json:"role"`
}
