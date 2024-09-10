package Config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	database, err := gorm.Open(
		postgres.Open("host=localhost port=5432 user=postgres dbname=bookGO sslmode=disable password=amir$$1379"),
		&gorm.Config{})

	if err != nil {
		panic(err)
	}

	db = database
}
func GetDB() *gorm.DB {
	return db
}
