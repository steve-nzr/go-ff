package database

import "go-ff/common/feature/inventory/def"

// Item database structure & table
type Item struct {
	ID           uint32 `gorm:"primary_key"`
	def.ItemBase `gorm:"embedded"`
	PlayerID     uint32
}
