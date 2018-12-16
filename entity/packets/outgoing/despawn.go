package outgoing

import (
	"go-ff/common/def/packet/snapshottype"
	"go-ff/common/service/cache"
	"go-ff/common/service/external"
)

// DeleteObj from the world
func DeleteObj(p *cache.Player) *external.Packet {
	return external.StartMergePacket(p.EntityID, snapshottype.Del_obj, 0xFFFFFF00)
}
