package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gofirebase/api"
)
func CreateDatabase() *gorm.DB {
	// Create db instance with gorm
	db, err := gorm.Open( "postgres", "host=localhost port=5432 user=postgres dbname=shaukat sslmode=disable password=root")
	if err != nil {
		panic("Failed to connect to database!")
	}
	// migrate our model for artist
	db.AutoMigrate(&api.Artist{})
	return db
}