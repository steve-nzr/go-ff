package main

import (
	"fmt"

	"flyff/cluster/packets"
	"flyff/core"
)

func main() {
	core.InitiateDbConnection()
	defer core.CloseDbConnection()

	db := core.GetDbConnection()
	db.AutoMigrate(&core.Character{})

	core.StartNetServer(core.NetServerConfig{
		Host:                "0.0.0.0",
		Port:                "28000",
		Type:                core.NetServerTCP,
		ConnectionInitiated: onConnectionInitiated,
		ConnectionMessage:   onConnectionMessage})
}

func onConnectionInitiated(c *core.NetClient) {
	clusterClient{c}.sendGreetings()
}

func onConnectionMessage(c *core.NetClient, packet *core.Packet) {
	cc := clusterClient{c}

	// Always FFFFFFF
	packet.ReadUInt32()

	protocol := packet.ReadUInt32()
	fmt.Printf("New packet with id : 0x%02x\n", protocol)

	if protocol == 0xf6 {
		cc.sendPlayerList(0).
			sendWorldAddr()
	} else if protocol == 0xf4 {
		var p packets.CreatePlayer
		p.Construct(packet)

		var c core.Character
		c.Slot = p.Slot
		c.Name = p.Name
		c.Gender = p.Gender
		c.MapID = 1
		c.PosX = 6968.0
		c.PosY = 100.0
		c.PosZ = 3328.0
		c.SkinSetID = uint32(p.SkinSet)
		c.HairID = uint32(p.HairMeshID)
		c.HairColor = p.HairColor
		c.FaceID = uint32(p.FaceID)
		c.JobID = p.Job
		c.Level = 1
		c.Strength = 15
		c.Stamina = 15
		c.Dexterity = 15
		c.Intelligence = 15

		db := core.GetDbConnection()
		db.Save(&c)

		cc.sendPlayerList(0)
	} else if protocol == 0xf5 {
		var p packets.DeletePlayer
		p.Construct(packet)

		db := core.GetDbConnection()
		db.Delete(&core.Character{}, p.CharacterID)

		cc.sendPlayerList(0)
	} else if protocol == 0xff05 {
		var p packets.PreJoin
		p.Construct(packet)

		cc.Send(core.MakePacket(core.PREJOIN))
	}
}
