package database

import (
	"github.com/Steve-Nzr/go-ff/pkg/feature/inventory/def"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	// Using with GORM Open
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Connection holds an active connection to the DB
var Connection *gorm.DB

// Initialize the DB connection
func Initialize() {
	db, err := gorm.Open("postgres", os.Getenv("POSTGRES_DB_ENDPOINT"))
	if err != nil {
		log.Fatal(err)
	}

	Connection = db
	autoMigrate()
}

func autoMigrate() {
	Connection.AutoMigrate(&Player{})
	Connection.AutoMigrate(&def.Item{})
}
