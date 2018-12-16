package outgoing

import (
	"go-ff/common/def/packet/snapshottype"
	"go-ff/common/service/cache"
	"go-ff/common/service/external"
)

func Chat(p *cache.Player, msg string) *external.Packet {
	return external.StartMergePacket(p.EntityID, snapshottype.Chat, 0xFFFFFF00).WriteString(msg).Finalize()
}
