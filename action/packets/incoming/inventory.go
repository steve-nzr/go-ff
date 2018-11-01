package incoming

import (
	"flyff/common/feature/inventory/helpers"
	"flyff/common/service/cache"
	"flyff/common/service/external"
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
