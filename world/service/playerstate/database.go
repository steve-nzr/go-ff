package playerstate

import (
	"flyff/world/entities"

	"github.com/jinzhu/gorm"

	// Using with GORM
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Connection *gorm.DB

func Initialize() {
	db, err := gorm.Open("mysql", "user:password@tcp(192.168.2.201:3306)/db?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	Connection = db
}

func AutoMigrate() {
	Connection.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&entities.PlayerEntity{})
}
