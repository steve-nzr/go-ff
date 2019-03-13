package outgoing

import (
	"go-ff/pkg/def/packet/snapshottype"
	"go-ff/pkg/service/cache"
	"go-ff/pkg/service/external"
)

func Chat(p *cache.Player, msg string) *external.Packet {
	return external.StartMergePacket(p.EntityID, snapshottype.Chat, 0xFFFFFF00).WriteString(msg).Finalize()
}
