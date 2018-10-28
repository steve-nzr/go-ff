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
	MaxItems        = 73
	EquipOffset     = MaxItems - 31
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
				ItemID:   -1,
			},
		})
	}

	for _, item := range items {
		if item.Position >= 73 {
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

func (i ItemContainer) Serialize(p *external.Packet) {
	var size uint8
	for index := 0; index < MaxItems; index++ {
		p.WriteInt32(i[index].UniqueID)
		if i[index].ItemID != -1 {
			size++
		}
	}

	p.WriteUInt8(size)

	for index := 0; index < MaxItems; index++ {
		if i[index].ItemID > 0 {
			p.WriteUInt8(uint8(i[index].UniqueID)).
				WriteInt32(i[index].UniqueID)
			i[index].Serialize(p)
		}
	}

	for index := 0; index < MaxItems; index++ {
		p.WriteInt32(i[index].UniqueID)
	}
}
