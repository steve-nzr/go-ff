package inventory

import (
	"flyff/common/feature/inventory/def"
	"flyff/common/service/database"
	"flyff/common/service/external"
)

// ItemContainer represents a list of items with the fixed size of an inventory
type ItemContainer []def.Item

const (
	RightWeaponSlot = 52
	EquipOffset     = 42
	MaxItems        = 73
	InventorySize   = EquipOffset
	MaxHumanParts   = MaxItems - EquipOffset
)

// InitializeWith database Item list
func (container ItemContainer) InitializeWith(items []database.Item) ItemContainer {
	for i := 0; i < MaxItems; i++ {
		container = append(container, def.Item{
			ItemBase: def.ItemBase{
				UniqueID: int32(i),
				Position: -1,
				Count:    -1,
			},
		})
	}

	for _, item := range items {
		if item.Position >= 73 {
			continue
		}

		container[item.Position] = def.Item{
			ItemBase: item.ItemBase,
		}
	}

	for i := EquipOffset; i < MaxItems; i++ {
		if container[i].ItemID == 0 {
			container[i].UniqueID = -1
		}
	}

	return container
}

func (i ItemContainer) Serialize(p *external.Packet) {
	var size uint8
	for _, item := range i {
		p.WriteInt32(item.UniqueID)
		if item.ItemID != 0 {
			size++
		}
	}

	p.WriteUInt8(size)

	for _, item := range i {
		if item.ItemID > 0 {
			p.WriteUInt8(uint8(item.UniqueID))
			p.WriteInt32(item.UniqueID)
			item.Serialize(p)
		}
	}

	for _, item := range i {
		p.WriteInt32(item.UniqueID)
	}
}
