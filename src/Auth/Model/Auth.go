package Model

import "github.com/google/uuid"

type Auth struct {
	Id       uuid.UUID `gorm:"type:uuid;default:null" json:"id"`
	UserName string    `gorm:"column:username;type:varchar(255);not null"`
	Email    string    `gorm:"column:email;type:varchar(255);not null"`
	Password string    `gorm:"column:password;type:varchar(255);not null"`
	Role     string    `gorm:"column:role;type:varchar(50);default:'User'"`
}
