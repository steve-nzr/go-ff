package core

import (
	"fmt"

	"github.com/jinzhu/gorm"

	// postgres to open the db
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var conn *gorm.DB

func setActiveConnection(db *gorm.DB) {
	conn = db
}

// InitiateDbConnection open a new db connection & set it to active
func InitiateDbConnection() {
	db, err := gorm.Open("postgres", "host=192.168.1.201 port=5432 user=user dbname=db password=password sslmode=disable")
	if err != nil {
		panic(err)
	}

	setActiveConnection(db)
	fmt.Println("[Database] Connection initiated")
}

// GetDbConnection returns the active db connection
func GetDbConnection() *gorm.DB {
	if conn == nil {
		panic("No active database connecion !")
	}

	return conn
}

// CloseDbConnection closes the active db connection
// To use with defer in your main progrem (e.g defer core.CloseConnection())
func CloseDbConnection() {
	if conn == nil {
		panic("No active database connecion !")
	}

	conn.Close()
	fmt.Println("[Database] Connection closed")
}
