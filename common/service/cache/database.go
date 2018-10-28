package cache

import (
	"flyff/common/feature/inventory/def"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	// Using with GORM Open
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Connection holds an active connection to the DB
var Connection *gorm.DB

// Initialize the Database connection & tables
func Initialize() {
	db, err := gorm.Open("mysql", os.Getenv("MYSQL_CACHE_ENDPOINT"))
	if err != nil {
		log.Fatal(err)
	}

	Connection = db
	autoMigrate()
}

func autoMigrate() {
	Connection.Set("gorm:table_options", "ENGINE=Memory").AutoMigrate(&Player{})
	Connection.Set("gorm:table_options", "ENGINE=Memory").AutoMigrate(&def.Item{})
}
