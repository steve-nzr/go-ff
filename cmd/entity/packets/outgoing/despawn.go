package outgoing

import (
	"go-ff/pkg/def/packet/snapshottype"
	"go-ff/pkg/service/cache"
	"go-ff/pkg/service/external"
)

// DeleteObj from the world
func DeleteObj(p *cache.Player) *external.Packet {
	return external.StartMergePacket(p.EntityID, snapshottype.Del_obj, 0xFFFFFF00)
}
