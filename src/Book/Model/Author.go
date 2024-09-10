package Model

import "github.com/google/uuid"

type Author struct {
	Id        uuid.UUID `gorm:"type:uuid;default:null" json:"id"`
	FirstName string    `gorm:"column:first_name" json:"firstName"`
	LastName  string    `gorm:"column:last_name" json:"lastName"`
}
