package def

import (
	"go-ff/common/service/external"
)

// Item holds in-game item data (not itemprops, not reference to a player)
type ItemBase struct {
	ItemID   int32
	UniqueID int32
	Count    int16
	Position int16
}

// Serialize the item to the packet
func (item ItemBase) Serialize(p *external.Packet) {
	p.WriteInt32(item.ItemID).
		WriteUInt32(0).
		WriteString("ItemName").
		WriteInt16(item.Count). // m_nItemNum
		WriteUInt8(0).          // m_nRepairNumber
		WriteUInt32(0).         // m_nHitPoint
		WriteUInt32(0).         // m_nRepair
		WriteUInt8(0).          // m_byFlag
		WriteUInt32(0).         // m_nAbilityOption
		WriteUInt32(0).         // m_idGuild
		WriteUInt8(0).          // m_bItemResist
		WriteUInt32(0).         // m_nResistAbilityOption
		WriteUInt32(0).         // m_nResistSMItemId
		WriteUInt32(0).         // m_vPiercing.size
		WriteUInt32(0).         // m_vUltimatePiercing.size
		WriteUInt32(0).         // SetVisKeepTime.size
		WriteUInt32(0).         // m_bCharged
		WriteUInt64(0).         // m_iRandomOptItemId
		WriteUInt32(0).         // m_dwKeepTime
		WriteUInt8(0).          // bPet
		WriteUInt32(0)          // m_bTranformVisPet
}
