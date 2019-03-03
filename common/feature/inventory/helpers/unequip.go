package helpers

import (
	"go-ff/common/feature/inventory"
	"go-ff/common/feature/inventory/def"
	"go-ff/common/feature/inventory/packets/outgoing"
	"go-ff/common/service/cache"
	"go-ff/common/service/external"
	"go-ff/common/service/messaging"
	"math"
)

func Equip(player *cache.Player, uniqueID uint8, part int8) {
	index := player.Inventory.GetItemIndex(uniqueID)
	if index < 0 {
		return
	}

	//item := player.Inventory[index]

	toEquip := part == -1
	if toEquip {
		/*if item.Position >= 0 {
			Unequip(player, uint8(item.UniqueID))
		}*/

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
	if item.Position <= inventory.EquipOffset {
		return false
	}

	availableSlot := player.Inventory.GetAvailableSlot()
	if availableSlot < 0 {
		// No space left in inventory
		return false
	}

	parts := int32(math.Abs(float64(item.Position - inventory.EquipOffset)))

	availableItem := &player.Inventory[availableSlot]
	item.ItemBase, availableItem.ItemBase = availableItem.ItemBase, item.ItemBase

	availableItem.Position = int16(availableSlot)
	item.UniqueID = -1
	item.Position = -1

	messaging.Publish(messaging.ConnectionTopic, &external.PacketEmitter{
		Packet: outgoing.Equip(player, availableItem, false, parts).Finalize(),
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

	sourceItem.Position = int16(destSlot)
	if destinationItem.Position != -1 {
		destinationItem.Position = int16(sourceSlot)
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
