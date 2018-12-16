package outgoing

import (
	"go-ff/common/def/packet/snapshottype"
	"go-ff/common/feature/inventory/def"
	"go-ff/common/service/cache"
	"go-ff/common/service/external"
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

func Update(player *cache.Player, item *def.Item, updatetype uint8, data int16, card uint16) *external.Packet {
	return external.StartMergePacket(player.EntityID, snapshottype.Moveitem, 0xFFFFFF00).
		WriteUInt8(0).
		WriteUInt8(uint8(item.UniqueID)).
		WriteUInt8(updatetype).
		WriteInt16(data).
		WriteUInt16(card)
}
