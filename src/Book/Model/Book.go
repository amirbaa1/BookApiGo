package Model

import "github.com/google/uuid"

type Book struct {
	Id        uuid.UUID `gorm:"type:uuid;default:null" json:"id"`
	Title     string    `json:"title"`
	AuthorID  uuid.UUID `json:"author_id"` // Foreign key
	Publisher string    `json:"publisher"`
	Author    Author    `gorm:"foreignKey:AuthorID" json:"author"`
}
