package outgoing

import (
	"flyff/common/def/packet/snapshottype"
	"flyff/common/feature/inventory/def"
	"flyff/common/service/cache"
	"flyff/common/service/external"
)

func Equip(player *cache.Player, item *def.Item, equip bool, targetSlot int32) *external.Packet {
	p := external.StartMergePacket(player.EntityID, snapshottype.Doequip, 0xFFFFFF00).
		WriteInt32(item.UniqueID).
		WriteUInt8(0)

	if equip {
		p.WriteInt8(1)
	} else {
		p.WriteInt8(0)
	}

	p.WriteInt32(item.ItemID).
		WriteUInt16(0).
		WriteUInt8(0).
		WriteUInt8(0).
		WriteUInt32(0).
		WriteInt32(targetSlot)

	return p
}
