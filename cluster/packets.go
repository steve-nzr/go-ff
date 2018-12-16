package main

import (
	"flyff/common/def/packet/packettype"
	"flyff/common/feature/inventory"
	"flyff/common/service/database"
	"flyff/common/service/external"

	. "github.com/ahmetb/go-linq"
)

func sendWorldAddr(c *external.Client) {
	packet := external.MakePacket(packettype.Cache_addr).
		WriteString("192.168.99.101")

	c.Send(packet)
}

func sendPlayerList(c *external.Client, authKey int32) {
	var characters []database.Player
	database.Connection.Limit(3).Preload("Items").Find(&characters)

	packet := external.MakePacket(packettype.Player_list).
		WriteInt32(0).
		WriteInt32(int32(len(characters)))

	for _, c := range characters {
		packet.WriteUInt32(uint32(c.Slot)).
			WriteUInt32(1).
			WriteUInt32(c.Position.MapID).
			WriteUInt32(0x0B + uint32(c.Gender)).
			WriteString(c.Name).
			WriteFloat32(float32(c.Position.Vec.X)).
			WriteFloat32(float32(c.Position.Vec.Y)).
			WriteFloat32(float32(c.Position.Vec.Z)).
			WriteUInt32(uint32(c.ID)).
			WriteUInt32(0).
			WriteUInt32(0).
			WriteUInt32(0).
			WriteUInt32(c.SkinSetID).
			WriteUInt32(c.HairID).
			WriteUInt32(c.HairColor).
			WriteUInt32(c.FaceID).
			WriteUInt8(c.Gender).
			WriteUInt32(uint32(c.JobID)).
			WriteUInt32(uint32(c.Level)).
			WriteUInt32(0).
			WriteUInt32(uint32(c.Statistics.Strength)).
			WriteUInt32(uint32(c.Statistics.Stamina)).
			WriteUInt32(uint32(c.Statistics.Dexterity)).
			WriteUInt32(uint32(c.Statistics.Intelligence)).
			WriteUInt32(0)

		var equipedItems []database.Item
		From(c.Items).
			Where(func(i interface{}) bool {
				item := i.(database.Item)
				return item.Position > inventory.EquipOffset
			}).
			Select(func(i interface{}) interface{} {
				return i.(database.Item)
			}).
			ToSlice(&equipedItems)

		packet.WriteUInt32(uint32(len(equipedItems)))
		for _, item := range equipedItems {
			packet.WriteInt32(item.ItemID)
		}
	}

	packet.WriteInt32(0)
	c.Send(packet)
}
