package outgoing

import (
	"flyff/common/service/cache"
	"flyff/common/service/external"
)

// DeleteObj from the world
func DeleteObj(p *cache.Player) *external.Packet {
	return external.StartMergePacket(p.EntityID, uint16(0x00F1), 0xFFFFFF00)
}
