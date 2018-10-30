package outgoing

import (
	"flyff/common/def/packet/snapshottype"
	"flyff/common/service/cache"
	"flyff/common/service/external"
)

func Chat(p *cache.Player, msg string) *external.Packet {
	return external.StartMergePacket(p.EntityID, snapshottype.Chat, 0xFFFFFF00).WriteString(msg).Finalize()
}
