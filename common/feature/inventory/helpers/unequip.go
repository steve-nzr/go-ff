package helpers

import (
	"flyff/common/feature/inventory"
	"flyff/common/feature/inventory/packets/outgoing"
	"flyff/common/service/cache"
	"flyff/common/service/external"
	"flyff/common/service/messaging"
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

	player.Inventory[sourceSlot], player.Inventory[destSlot] = player.Inventory[destSlot], player.Inventory[sourceSlot]
	messaging.Publish(messaging.ConnectionTopic, &external.PacketEmitter{
		Packet: outgoing.Move(player, sourceSlot, destSlot).Finalize(),
		To:     cache.FindIDAround(player),
	})
}
