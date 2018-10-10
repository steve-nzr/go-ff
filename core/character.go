package core

import (
	"github.com/jinzhu/gorm"
)

// Vector3D is a 3 dimensional vector/point
type Vector3D struct {
	x float32
	y float32
	z float32
}

// Character holds a complete account's character/player
type Character struct {
	gorm.Model
	Slot         uint8  `gorm:"type:SMALLINT"`
	Name         string `gorm:"type:VARCHAR(32)"`
	Gender       uint8  `gorm:"type:SMALLINT"`
	MapID        uint32 `gorm:"type:SMALLINT"`
	PosX         float32
	PosY         float32
	PosZ         float32
	SkinSetID    uint32
	HairID       uint32
	HairColor    uint32
	FaceID       uint32
	JobID        uint8
	Level        uint16
	Strength     uint16
	Stamina      uint16
	Dexterity    uint16
	Intelligence uint16
}
