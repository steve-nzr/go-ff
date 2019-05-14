package incoming

import (
	"go-ff/common/feature/inventory/helpers"
	"go-ff/common/service/cache"
	"go-ff/common/service/external"
)

// EquipItem packet
func EquipItem(p *external.PacketHandler) {
	player := cache.FindByNetID(p.ClientID)
	if player == nil {
		return
	}

	uniqueID := p.Packet.ReadUInt32()
	part := p.Packet.ReadInt32()

	helpers.Equip(player, (uint8)(uniqueID), part)
}

// MoveItem packet
func MoveItem(p *external.PacketHandler) {
	player := cache.FindByNetID(p.ClientID)
	if player == nil {
		return
	}

	p.Packet.ReadUInt8() // skipped
	sourceSlot := p.Packet.ReadUInt8()
	destSlot := p.Packet.ReadUInt8()

	helpers.Move(player, sourceSlot, destSlot)
}

func DropItem(p *external.PacketHandler) {
	player := cache.FindByNetID(p.ClientID)
	if player == nil {
		return
	}

	p.Packet.ReadUInt32()
	uniqueID := p.Packet.ReadUInt32()
	count := p.Packet.ReadInt16()
	helpers.Drop(player, uniqueID, count)
}
