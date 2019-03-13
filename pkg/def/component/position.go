package component

import "github.com/golang/geo/r3"

// Position Component (Including MapID)
type Position struct {
	MapID uint32
	Vec   r3.Vector `gorm:"embedded"`
}
