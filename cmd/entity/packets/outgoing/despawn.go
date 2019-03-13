package outgoing

import (
	"github.com/Steve-Nzr/go-ff/pkg/def/packet/snapshottype"
	"github.com/Steve-Nzr/go-ff/pkg/service/cache"
	"github.com/Steve-Nzr/go-ff/pkg/service/external"
)

// DeleteObj from the world
func DeleteObj(p *cache.Player) *external.Packet {
	return external.StartMergePacket(p.EntityID, snapshottype.Del_obj, 0xFFFFFF00)
}
