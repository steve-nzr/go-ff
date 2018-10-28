package def

import (
	"flyff/common/service/external"
)

// Item holds in-game item data (not itemprops, not reference to a player)
type ItemBase struct {
	ItemID   uint32
	UniqueID int32
	Count    int16
	Position int16
}

// Serialize the item to the packet
func (item ItemBase) Serialize(p *external.Packet) {
	p.WriteUInt32(item.ItemID).
		WriteUInt32(0).
		WriteUInt32(31).
		WriteInt16(item.Count).
		WriteUInt8(0).
		WriteUInt32(0).
		WriteUInt32(0).
		WriteUInt8(0).
		WriteUInt32(0).
		WriteUInt32(0).
		WriteUInt8(0).
		WriteUInt32(0).
		WriteUInt32(0).
		WriteUInt32(0).
		WriteUInt32(0).
		WriteUInt32(0).
		WriteUInt64(0).
		WriteUInt32(0).
		WriteUInt8(0).
		WriteUInt32(0)
}
