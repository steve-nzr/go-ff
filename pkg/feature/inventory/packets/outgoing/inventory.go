package outgoing

import (
	"github.com/Steve-Nzr/go-ff/pkg/def/packet/snapshottype"
	"github.com/Steve-Nzr/go-ff/pkg/feature/inventory/def"
	"github.com/Steve-Nzr/go-ff/pkg/service/cache"
	"github.com/Steve-Nzr/go-ff/pkg/service/external"
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

func Move(player *cache.Player, sourceSlot uint8, destSlot uint8) *external.Packet {
	return external.StartMergePacket(player.EntityID, snapshottype.Moveitem, 0xFFFFFF00).
		WriteUInt8(0).
		WriteUInt8(sourceSlot).
		WriteUInt8(destSlot)
}

func Update(player *cache.Player, uniqueID int32, updatetype uint8, data int32) *external.Packet {
	return external.StartMergePacket(player.EntityID, snapshottype.Update_item, 0xFFFFFF00).
		WriteUInt8(0).
		WriteUInt8(uint8(uniqueID)).
		WriteUInt8(updatetype).
		WriteInt32(data).
		WriteUInt32(0)
}
