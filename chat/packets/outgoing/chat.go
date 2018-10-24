package outgoing

import (
	"flyff/common/service/cache"
	"flyff/common/service/external"
)

func Chat(p *cache.Player, msg string) *external.Packet {
	return external.StartMergePacket(p.EntityID, 0x0001, 0xFFFFFF00).WriteString(msg).Finalize()
}
