package helpers

import (
	"go-ff/common/feature/inventory"
	"go-ff/common/feature/inventory/def"
	"go-ff/common/feature/inventory/packets/outgoing"
	"go-ff/common/service/cache"
	"go-ff/common/service/external"
	"go-ff/common/service/messaging"
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

	cache.Connection.Save(&player.Inventory[index])
	cache.Connection.Save(&player.Inventory[availableSlot])
	return true
}

func Move(player *cache.Player, sourceSlot uint8, destSlot uint8) {
	if sourceSlot > inventory.EquipOffset ||
		destSlot > inventory.EquipOffset {
		return
	}

	player.Inventory[sourceSlot].Position = int16(destSlot)

	if player.Inventory[destSlot].Position != -1 {
		player.Inventory[destSlot].Position = int16(sourceSlot)
	}

	player.Inventory[sourceSlot].ItemBase, player.Inventory[destSlot].ItemBase = player.Inventory[destSlot].ItemBase, player.Inventory[sourceSlot].ItemBase

	messaging.Publish(messaging.ConnectionTopic, &external.PacketEmitter{
		Packet: outgoing.Move(player, sourceSlot, destSlot).Finalize(),
		To:     []uint32{player.NetClientID},
	})

	cache.Connection.Save(&player.Inventory[sourceSlot])
	cache.Connection.Save(&player.Inventory[destSlot])
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
		To:     []uint32{player.NetClientID},
	})
}
