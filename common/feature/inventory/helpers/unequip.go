package helpers

import (
	"flyff/common/feature/inventory"
	"flyff/common/feature/inventory/def"
	"flyff/common/feature/inventory/packets/outgoing"
	"flyff/common/service/cache"
	"flyff/common/service/external"
	"flyff/common/service/messaging"
	"fmt"
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
	item := player.Inventory[index]
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

	availableItem := player.Inventory[availableSlot]
	item.Position = int16(availableSlot)
	player.Inventory[index], player.Inventory[availableSlot] = availableItem, item

	messaging.Publish(messaging.ConnectionTopic, &external.PacketEmitter{
		Packet: outgoing.Equip(player, &item, false, parts).Finalize(),
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
	sourceItem.Position = int16(destSlot)

	destItem := &player.Inventory[destSlot]
	if destItem.Position != -1 {
		destItem.Position = int16(sourceSlot)
	}

	player.Inventory[sourceSlot], player.Inventory[destSlot] = *destItem, *sourceItem

	messaging.Publish(messaging.ConnectionTopic, &external.PacketEmitter{
		Packet: outgoing.Move(player, sourceSlot, destSlot).Finalize(),
		To:     cache.FindIDAround(player),
	})

	cache.Connection.Model(&player).Association("Inventory").Replace(player.Inventory)
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

	item := player.Inventory[index]
	fmt.Println(item)
	if count >= item.Count {
		player.Inventory[index] = def.Item{
			ItemBase: def.ItemBase{
				UniqueID: int32(index),
				Position: -1,
				Count:    -1,
				ItemID:   -1,
			},
		}
	} else {
		item.Count -= count
		player.Inventory[index] = item
	}

	messaging.Publish(messaging.ConnectionTopic, &external.PacketEmitter{
		Packet: outgoing.Update(player, &item, def.ItmUpdateCount, item.Count, 0).Finalize(),
		To:     cache.FindIDAround(player),
	})
}
