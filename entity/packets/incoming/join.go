package incoming

import (
	"flyff/common/service/cache"
	"flyff/common/service/database"
	"flyff/common/service/external"
	"flyff/common/service/messaging"
	"flyff/entity/packets/outgoing"
	"log"
)

// JoinPacket struct
type JoinPacket struct {
	WorldID           uint32
	PlayerID          uint32
	AuthenticationKey int32
	PartyID           uint32
	GuildID           uint32
	GuildWarID        uint32
	IDOfMulti         int32
	Slot              uint8
	PlayerName        string
	Username          string
	Password          string
	MessengerState    int32
	MessengerCount    int32
}

// Construct ...
func (d *JoinPacket) Construct(p *external.Packet) {
	d.WorldID = p.ReadUInt32()
	d.PlayerID = p.ReadUInt32()
	d.AuthenticationKey = p.ReadInt32()
	d.PartyID = p.ReadUInt32()
	d.GuildID = p.ReadUInt32()
	d.GuildWarID = p.ReadUInt32()
	d.IDOfMulti = p.ReadInt32()
	d.Slot = p.ReadUInt8()
	d.PlayerName = p.ReadString()
	d.Username = p.ReadString()
	d.Password = p.ReadString()
	d.MessengerState = p.ReadInt32()
	d.MessengerCount = p.ReadInt32()
}

// Join World request
func Join(p *external.PacketHandler) {
	var join JoinPacket
	join.Construct(p.Packet)

	db := database.Connection

	var player database.Player
	player.ID = join.PlayerID

	res := db.First(&player)
	if res.Error != nil {
		log.Print(res.Error)
		return
	}

	var entitiy cache.Player
	entitiy.NetClientID = p.ClientID
	entitiy.EntityID = external.GenerateID()
	entitiy.PlayerID = uint32(player.ID)
	entitiy.Slot = player.Slot
	entitiy.HairColor = player.HairColor
	entitiy.HairID = player.HairID
	entitiy.FaceID = player.FaceID
	entitiy.SkinSetID = player.SkinSetID
	entitiy.JobID = player.JobID
	entitiy.Angle = 360.0
	entitiy.Gender = player.Gender
	entitiy.Level = player.Level
	entitiy.Type = 5 // Mover
	entitiy.Size = 100
	entitiy.Position = player.Position
	entitiy.Name = player.Name
	if player.Gender == 0 {
		entitiy.ModelID = 11
	} else if player.Gender == 1 {
		entitiy.ModelID = 12
	}
	entitiy.Statistics = player.Statistics

	// Tx BEGIN ----
	tx := cache.Connection.Begin()
	res = tx.Save(&entitiy)
	if res.Error != nil {
		log.Print(res.Error)
		tx.Rollback()
		return
	}

	tx.Commit()
	// Tx END ----

	messaging.Publish(messaging.ConnectionTopic, &external.PacketEmitter{
		Packet: outgoing.Spawn(&entitiy),
		To:     []uint32{p.ClientID},
	})
}
