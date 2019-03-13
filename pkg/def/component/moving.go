package component

import "github.com/golang/geo/r3"

// Moving component represents the player's move behaviour
type Moving struct {
	Destination r3.Vector `gorm:"embedded"`
	Motion      uint32
	Angle       float64
}
