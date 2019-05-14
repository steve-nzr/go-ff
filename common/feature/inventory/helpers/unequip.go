package helpers

import (
	"go-ff/common/feature/inventory"
	"go-ff/common/feature/inventory/def"
	"go-ff/common/feature/inventory/packets/outgoing"
	"go-ff/common/service/cache"
	"go-ff/common/service/external"
	"go-ff/common/service/messaging"
	"go-ff/common/service/resources"
)

func Equip(player *cache.Player, uniqueID uint8, part int32) {
	index := player.Inventory.GetItemIndex(uniqueID)
	if index < 0 {
		return
	}

	toEquip := part == -1
	if toEquip {
		prop := resources.ItemsProp[player.Inventory[index].ItemID]

		equipedItemSlot := prop.Parts + inventory.EquipOffset
		equipedItem := &player.Inventory[equipedItemSlot]
		if equipedItem.ItemID != -1 {
			Unequip(player, (uint8)(equipedItem.UniqueID))
		}

		index := player.Inventory.GetItemIndex(uniqueID) // add check
		item := &player.Inventory[index]

		item.ItemBase, equipedItem.ItemBase = equipedItem.ItemBase, item.ItemBase
		equipedItem.Position = (int16)(equipedItemSlot)

		messaging.Publish(messaging.ConnectionTopic, &external.PacketEmitter{
			Packet: outgoing.Equip(player, equipedItem, true, (int32)(prop.Parts)).Finalize(),
			To:     cache.FindIDAround(player),
		})

	} else {
		Unequip(player, uniqueID)
	}

	cache.Connection.Model(&player).Association("Inventory").Replace(player.Inventory)
}

func Unequip(player *cache.Player, uniqueID uint8) bool {
	index := player.Inventory.GetItemIndex(uniqueID)
	if index < 0 {
		return false
	}

	item := &player.Inventory[index]
	if item.Position <= inventory.EquipOffset {
		return false
	}

	availableSlot := player.Inventory.GetAvailableSlot()
	if availableSlot < 0 {
		return false
	}

	targetItem := &player.Inventory[availableSlot]
	parts := resources.ItemsProp[item.ItemID].Parts

	item.ItemBase, targetItem.ItemBase = targetItem.ItemBase, item.ItemBase
	targetItem.Position = (int16)(availableSlot)

	messaging.Publish(messaging.ConnectionTopic, &external.PacketEmitter{
		Packet: outgoing.Equip(player, targetItem, false, (int32)(parts)).Finalize(),
		To:     cache.FindIDAround(player),
	})

	return true
}

func Move(player *cache.Player, sourceSlot uint8, destSlot uint8) {
	if sourceSlot > inventory.EquipOffset ||
		destSlot > inventory.EquipOffset {
		return
	}

	sourceItem := &player.Inventory[sourceSlot]
	destinationItem := &player.Inventory[destSlot]
	if sourceItem.Count < 1 {
		return
	}

	sourceItem.Position = (int16)(destSlot)
	if destinationItem.Position != -1 {
		destinationItem.Position = (int16)(sourceSlot)
	}

	sourceItem.ItemBase, destinationItem.ItemBase = destinationItem.ItemBase, sourceItem.ItemBase

	messaging.Publish(messaging.ConnectionTopic, &external.PacketEmitter{
		Packet: outgoing.Move(player, sourceSlot, destSlot).Finalize(),
		To:     []uint32{player.NetClientID},
	})

	cache.Connection.Save(sourceItem)
	cache.Connection.Save(destinationItem)
}

func Drop(player *cache.Player, uniqueID uint32, count int16) {
	if uniqueID > inventory.MaxItems ||
		count < 1 {
		return
	}

	index := player.Inventory.GetItemIndex(uint8(uniqueID))
	if index < 0 {
		return
	}

	item := &player.Inventory[index]
	if count >= item.Count {
		item.ItemBase = def.ItemBase{
			ItemID:   -1,
			UniqueID: item.UniqueID,
			Count:    -1,
			Position: -1,
		}
	} else {
		item.Count -= count
	}

	messaging.Publish(messaging.ConnectionTopic, &external.PacketEmitter{
		Packet: outgoing.Update(player, item.UniqueID, def.ItmUpdateCount, (int32)(item.Count)).Finalize(),
		To:     []uint32{player.NetClientID},
	})

	cache.Connection.Save(item)
}
