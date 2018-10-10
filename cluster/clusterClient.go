package main

import (
	"flyff/core"
)

type clusterClient struct {
	*core.NetClient
}

func (nc clusterClient) sendGreetings() clusterClient {
	p := core.MakePacket(core.GREETINGS).
		WriteUInt32(nc.ID)

	nc.Send(p)
	return nc
}

func (nc clusterClient) sendWorldAddr() clusterClient {
	packet := core.MakePacket(core.WORLDADDR).
		WriteString("127.0.0.1")

	nc.Send(packet)
	return nc
}

func (nc clusterClient) sendPlayerList(authKey int32) clusterClient {
	db := core.GetDbConnection()

	var characters []core.Character
	db.Limit(3).Find(&characters)

	packet := core.MakePacket(core.PLAYERLIST).
		WriteInt32(0).
		WriteInt32(int32(len(characters)))

	for _, c := range characters {
		packet = packet.WriteUInt32(uint32(c.Slot)).
			WriteUInt32(1).
			WriteUInt32(c.MapID).
			WriteUInt32(0x0B + uint32(c.Gender)).
			WriteString(c.Name).
			WriteFloat32(c.PosX).
			WriteFloat32(c.PosY).
			WriteFloat32(c.PosZ).
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
			WriteUInt32(uint32(c.Strength)).
			WriteUInt32(uint32(c.Stamina)).
			WriteUInt32(uint32(c.Dexterity)).
			WriteUInt32(uint32(c.Intelligence)).
			WriteUInt32(0).
			WriteUInt32(0)

		/*for i := 0; i < 31; i++ {
			packet = packet.WriteInt32(-1)
		}*/
	}

	packet = packet.WriteInt32(0)

	nc.Send(packet)
	return nc
}
