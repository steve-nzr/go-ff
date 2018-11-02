package incoming

import (
	"flyff/common/feature/inventory/helpers"
	"flyff/common/service/cache"
	"flyff/common/service/external"
	"fmt"
)

// EquipItem packet
func EquipItem(p *external.PacketHandler) {
	player := cache.FindByNetID(p.ClientID)
	if player == nil {
		return
	}

	uniqueID := p.Packet.ReadUInt8()
	part := p.Packet.ReadInt8()

	helpers.Equip(player, uniqueID, part)
}

// MoveItem packet
func MoveItem(p *external.PacketHandler) {
	fmt.Println("Begin move")
	player := cache.FindByNetID(p.ClientID)
	if player == nil {
		return
	}

	p.Packet.ReadUInt8() // skipped
	sourceSlot := p.Packet.ReadUInt8()
	destSlot := p.Packet.ReadUInt8()

	helpers.Move(player, sourceSlot, destSlot)
	fmt.Println("End move")
}

func DropItem(p *external.PacketHandler) {
	fmt.Println("Begin Drop")
	player := cache.FindByNetID(p.ClientID)
	if player == nil {
		return
	}

	p.Packet.ReadUInt32()
	uniqueID := p.Packet.ReadUInt32()
	count := p.Packet.ReadInt16()
	helpers.Drop(player, uniqueID, count)
	fmt.Println("End Drop")
}
