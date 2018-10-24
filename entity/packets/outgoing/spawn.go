package outgoing

import (
	"flyff/common/service/cache"
	"flyff/common/service/external"
	"math"
)

// Spawn packet
func Spawn(pe *cache.Player) *external.Packet {
	p := external.StartMergePacket(uint32(pe.EntityID), uint16(external.ENVIRONMENTALL), 0x0000FF00).
		WriteUInt32(0)

	p.AddMergePart(external.WORLDREADINFO, uint32(pe.EntityID)).
		WriteUInt32(pe.Position.MapID).
		WriteFloat32(float32(pe.Position.Vec.X)).
		WriteFloat32(float32(pe.Position.Vec.Y)).
		WriteFloat32(float32(pe.Position.Vec.Z))

	p.AddMergePart(external.ADDOBJ, uint32(pe.EntityID))

	p.WriteUInt8(5)
	if pe.Gender == 0 {
		p.WriteUInt32(11)
	} else if pe.Gender == 1 {
		p.WriteUInt32(12)
	}

	p.WriteUInt8(5)
	if pe.Gender == 0 {
		p.WriteUInt32(11)
	} else if pe.Gender == 1 {
		p.WriteUInt32(12)
	}

	p.WriteUInt16(100).
		WriteFloat32(float32(pe.Position.Vec.X)).
		WriteFloat32(float32(pe.Position.Vec.Y)).
		WriteFloat32(float32(pe.Position.Vec.Z)).
		WriteUInt16(360).
		WriteUInt32(uint32(pe.EntityID))

	p.WriteUInt16(0).
		WriteUInt8(1).
		WriteUInt32(230).
		WriteUInt32(0).
		WriteUInt32(0).
		WriteUInt8(1).
		WriteInt32(-1)

	p.WriteString(pe.Name).
		WriteUInt8(pe.Gender).
		WriteUInt8(uint8(pe.SkinSetID)).
		WriteUInt8(uint8(pe.HairID)).
		WriteUInt32(pe.HairColor).
		WriteUInt8(uint8(pe.FaceID)).
		WriteUInt32(uint32(pe.PlayerID)). // Playerdata ID
		WriteUInt8(uint8(pe.JobID)).
		WriteUInt16(pe.Statistics.Strength).
		WriteUInt16(pe.Statistics.Stamina).
		WriteUInt16(pe.Statistics.Dexterity).
		WriteUInt16(pe.Statistics.Intelligence).
		WriteUInt16(uint16(pe.Level))

	p.WriteInt32(-1).
		WriteUInt32(0).
		WriteUInt8(0). // Have guild or not
		WriteUInt32(0).
		WriteUInt8(0) // Have party

	// Auth
	p.WriteUInt8(100).
		WriteUInt32(0).
		WriteUInt32(0).
		WriteUInt32(0x000001F6). // item used ??
		WriteUInt32(0).
		WriteUInt32(0).
		WriteUInt32(0).
		WriteUInt32(0).
		WriteUInt32(0).
		WriteUInt8(0). // duel
		WriteInt32(-1) // titles

	// 31 = MaxItems - EquipOffset
	for i := 0; i < 31; i++ {
		p.WriteUInt32(0)
	}

	p.WriteUInt32(0)

	for i := 0; i < 26; i++ {
		p.WriteUInt32(0)
	}

	p.WriteUInt16(0).
		WriteUInt16(0).
		WriteUInt32(0).
		WriteUInt32(0).
		WriteUInt32(0). // Gold
		WriteUInt64(0). // exp
		WriteUInt32(0).
		WriteUInt32(0).
		WriteUInt64(0). // Death exp
		WriteUInt32(0)

	for i := 0; i < 32; i++ {
		p.WriteUInt32(0)
	}

	// Marking pos
	p.WriteUInt32(pe.Position.MapID).
		WriteFloat32(float32(pe.Position.Vec.X)).
		WriteFloat32(float32(pe.Position.Vec.Y)).
		WriteFloat32(float32(pe.Position.Vec.Z))

	p.WriteUInt8(0).
		WriteUInt8(0).
		WriteUInt8(0)

	p.WriteUInt32(42).
		WriteUInt16(0). // Stats points
		WriteUInt16(0)

	for i := 0; i < 31; i++ {
		p.WriteInt32(-1)
	}

	for i := 0; i < 45; i++ {
		p.WriteInt32(-1).
			WriteUInt32(0)
	}

	p.WriteUInt8(0). // Cheer point
				WriteUInt32(0)

	p.WriteUInt8(pe.Slot)
	for i := 0; i < 3; i++ {
		p.WriteUInt32(0).
			WriteUInt32(0)
	}

	p.WriteUInt32(1).
		WriteUInt8(0).
		WriteUInt8(0).
		WriteUInt8(0).
		WriteUInt64(0).
		WriteUInt32(0)

	// Serialize inventory
	for i := 0; i < 73; i++ {
		p.WriteInt32(-1)
	}
	p.WriteUInt8(0) // Item count
	for i := 0; i < 73; i++ {
		p.WriteInt32(-1)
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 42; j++ {
			p.WriteUInt32(uint32(j))
		}
		p.WriteUInt8(0)
		for j := 0; j < 42; j++ {
			p.WriteUInt32(uint32(j))
		}
	}

	p.WriteInt32(math.MaxInt32)

	// Bag
	p.WriteInt8(1)
	for i := 0; i < 6; i++ {
		p.WriteUInt32(uint32(i))
	}
	p.WriteInt8(0)
	for i := 0; i < 6; i++ {
		p.WriteUInt32(uint32(i))
	}

	p.WriteUInt32(1).
		WriteUInt32(1).
		WriteInt8(0).
		WriteUInt32(1)

	for i := 0; i < 150; i++ {
		p.WriteUInt32(0)
	}

	p.WriteUInt32(0).
		WriteUInt32(0).
		WriteUInt32(0)

	return p
}
