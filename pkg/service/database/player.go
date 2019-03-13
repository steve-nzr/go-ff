package database

import (
	"go-ff/pkg/def/component"

	"github.com/jinzhu/gorm"
)

// Player holds a complete account's character/player
type Player struct {
	ID         uint32 `gorm:"primary_key"`
	Slot       uint8
	Name       string `gorm:"type:VARCHAR(32)"`
	Gender     uint8
	SkinSetID  uint32
	HairID     uint32
	HairColor  uint32
	FaceID     uint32
	JobID      uint8
	Level      uint16
	Position   component.Position   `gorm:"embedded;EMBEDDED_PREFIX:posit_"`
	Statistics component.Statistics `gorm:"embedded;EMBEDDED_PREFIX:stats_"`
	Items      []Item
}

// BeforeDelete the player
func (p *Player) BeforeDelete(scope *gorm.Scope) error {
	Connection.Delete(&p.Items)
	return nil
}
