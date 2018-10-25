package component

import "github.com/golang/geo/r3"

// Moving component represents the player's destination
type Moving struct {
	Vec r3.Vector `gorm:"embedded"`
}
