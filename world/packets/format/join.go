package format

import "flyff/core"

// Join packet struct
type Join struct {
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
func (d *Join) Construct(p *core.Packet) {
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
