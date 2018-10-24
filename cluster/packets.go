package main

import (
	"flyff/common/service/database"
	"flyff/common/service/external"
)

func sendWorldAddr(c *external.Client) {
	packet := external.MakePacket(external.WORLDADDR).
		WriteString("127.0.0.1")

	c.Send(packet)
}

func sendPlayerList(c *external.Client, authKey int32) {
	var characters []database.Player
	database.Connection.Limit(3).Find(&characters)

	packet := external.MakePacket(external.PLAYERLIST).
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
			WriteUInt32(0).
			WriteUInt32(0)
	}

	packet.WriteInt32(0)
	c.Send(packet)
}
