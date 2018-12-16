package outgoing

import (
	"flyff/common/def/packet/snapshottype"
	"flyff/common/service/cache"
	"flyff/common/service/external"
)

// DeleteObj from the world
func DeleteObj(p *cache.Player) *external.Packet {
	return external.StartMergePacket(p.EntityID, snapshottype.Del_obj, 0xFFFFFF00)
}
