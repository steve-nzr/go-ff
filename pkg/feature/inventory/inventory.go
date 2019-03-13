package inventory

import (
	"go-ff/pkg/feature/inventory/def"
	"go-ff/pkg/service/database"
	"go-ff/pkg/service/external"
)

// ItemContainer represents a list of items with the fixed size of an inventory
type ItemContainer []def.Item

const (
	RightWeaponSlot = 52
	MaxItems        = 73
	EquipOffset     = MaxItems - 31
	InventorySize   = EquipOffset
	MaxHumanParts   = MaxItems - EquipOffset
)

func (container ItemContainer) GetItemIndex(uniqueID uint8) int32 {
	for i, item := range container {
		if item.UniqueID == int32(uniqueID) {
			return int32(i)
		}
	}

	return -1
}

func (container ItemContainer) GetAvailableSlot() int32 {
	for i, item := range container {
		if item.Position == -1 {
			return int32(i)
		}
	}

	return -1
}

// InitializeWith database Item list
func (container ItemContainer) InitializeWith(items []database.Item) ItemContainer {
	for i := 0; i < MaxItems; i++ {
		container = append(container, def.Item{
			ItemBase: def.ItemBase{
				UniqueID: int32(i),
				Position: -1,
				Count:    -1,
				ItemID:   -1,
			},
		})
	}

	for _, item := range items {
		if item.Position >= MaxItems || item.Position < 0 {
			continue
		}

		item.UniqueID = container[item.Position].UniqueID
		container[item.Position] = def.Item{
			ItemBase: item.ItemBase,
		}
	}

	for i := EquipOffset; i < MaxItems; i++ {
		if container[i].ItemID == -1 {
			container[i].UniqueID = -1
		}
	}

	return container
}

// ConvertToDatabaseSlice creates a slice to be savec to the persistent database
func (container ItemContainer) ConvertToDatabaseSlice() []database.Item {
	var items []database.Item
	for _, item := range container {
		if item.ItemID < 1 {
			continue
		}

		items = append(items, database.Item{
			PlayerID: item.PlayerID,
			ItemBase: item.ItemBase,
		})
	}

	return items
}

// Serialize the entire inventory to the given packet
func (container ItemContainer) Serialize(p *external.Packet) {
	var size uint8
	for i := 0; i < MaxItems; i++ {
		p.WriteInt32(container[i].UniqueID)
		if container[i].ItemID != -1 {
			size++
		}
	}

	p.WriteUInt8(size)

	for i := 0; i < MaxItems; i++ {
		if container[i].ItemID > 0 {
			p.WriteUInt8(uint8(container[i].UniqueID)).
				WriteInt32(container[i].UniqueID)
			container[i].Serialize(p)
		}
	}

	for i := 0; i < MaxItems; i++ {
		p.WriteInt32(container[i].UniqueID)
	}
}
