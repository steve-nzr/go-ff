package def

// Item cache structure & table
type Item struct {
	ID       uint32 `gorm:"primary_key"`
	ItemBase `gorm:"embedded"`
	PlayerID uint32
}
