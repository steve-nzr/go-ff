package out

import (
	"flyff/core/net"
	"flyff/world/game/structure"
	"math"
)

func MakeSpawn(from *structure.PlayerEntity) net.Packet {
	p := net.StartMergePacket(uint32(from.ID), uint16(net.ENVIRONMENTALL), 0x0000FF00).
		WriteUInt32(0)

	p = p.AddMergePart(net.WORLDREADINFO, uint32(from.ID)).
		WriteUInt32(from.Position.MapID).
		WriteFloat32(float32(from.Position.Vec.X)).
		WriteFloat32(float32(from.Position.Vec.Y)).
		WriteFloat32(float32(from.Position.Vec.Z))

	p = p.AddMergePart(net.ADDOBJ, uint32(from.ID))

	p = p.WriteUInt8(5)
	if from.Gender == 0 {
		p = p.WriteUInt32(11)
	} else if from.Gender == 1 {
		p = p.WriteUInt32(12)
	}

	p = p.WriteUInt8(5)
	if from.Gender == 0 {
		p = p.WriteUInt32(11)
	} else if from.Gender == 1 {
		p = p.WriteUInt32(12)
	}

	p = p.WriteUInt16(100).
		WriteFloat32(float32(from.Position.Vec.X)).
		WriteFloat32(float32(from.Position.Vec.Y)).
		WriteFloat32(float32(from.Position.Vec.Z)).
		WriteUInt16(360).
		WriteUInt32(uint32(from.ID))

	p = p.WriteUInt16(0).
		WriteUInt8(1).
		WriteUInt32(230).
		WriteUInt32(0).
		WriteUInt32(0).
		WriteUInt8(1).
		WriteInt32(-1)

	p = p.WriteString(from.Name).
		WriteUInt8(from.Gender).
		WriteUInt8(uint8(from.SkinSetID)).
		WriteUInt8(uint8(from.HairID)).
		WriteUInt32(from.HairColor).
		WriteUInt8(uint8(from.FaceID)).
		WriteUInt32(uint32(from.PlayerID)). // Playerdata ID
		WriteUInt8(uint8(from.JobID)).
		WriteUInt16(from.Statistics.Strength).
		WriteUInt16(from.Statistics.Stamina).
		WriteUInt16(from.Statistics.Dexterity).
		WriteUInt16(from.Statistics.Intelligence).
		WriteUInt16(uint16(from.Level))

	p = p.WriteInt32(-1).
		WriteUInt32(0).
		WriteUInt8(0). // Have guild or not
		WriteUInt32(0).
		WriteUInt8(0) // Have party

	// Auth
	p = p.WriteUInt8(100).
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
		p = p.WriteUInt32(0)
	}

	p = p.WriteUInt32(0)

	for i := 0; i < 26; i++ {
		p = p.WriteUInt32(0)
	}

	p = p.WriteUInt16(0).
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
		p = p.WriteUInt32(0)
	}

	// Marking pos
	p = p.WriteUInt32(from.Position.MapID).
		WriteFloat32(float32(from.Position.Vec.X)).
		WriteFloat32(float32(from.Position.Vec.Y)).
		WriteFloat32(float32(from.Position.Vec.Z))

	p = p.WriteUInt8(0).
		WriteUInt8(0).
		WriteUInt8(0)

	p = p.WriteUInt32(42).
		WriteUInt16(0). // Stats points
		WriteUInt16(0)

	for i := 0; i < 31; i++ {
		p = p.WriteInt32(-1)
	}

	for i := 0; i < 45; i++ {
		p = p.WriteInt32(-1).
			WriteUInt32(0)
	}

	p = p.WriteUInt8(0). // Cheer point
				WriteUInt32(0)

	p = p.WriteUInt8(from.Slot)
	for i := 0; i < 3; i++ {
		p = p.WriteUInt32(0).
			WriteUInt32(0)
	}

	p = p.WriteUInt32(1).
		WriteUInt8(0).
		WriteUInt8(0).
		WriteUInt8(0).
		WriteUInt64(0).
		WriteUInt32(0)

	// Serialize inventory
	for i := 0; i < 73; i++ {
		p = p.WriteInt32(-1)
	}
	p = p.WriteUInt8(0) // Item count
	for i := 0; i < 73; i++ {
		p = p.WriteInt32(-1)
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 42; j++ {
			p = p.WriteUInt32(uint32(j))
		}
		p = p.WriteUInt8(0)
		for j := 0; j < 42; j++ {
			p = p.WriteUInt32(uint32(j))
		}
	}

	p = p.WriteInt32(math.MaxInt32)

	// Bag
	p = p.WriteInt8(1)
	for i := 0; i < 6; i++ {
		p = p.WriteUInt32(uint32(i))
	}
	p = p.WriteInt8(0)
	for i := 0; i < 6; i++ {
		p = p.WriteUInt32(uint32(i))
	}

	p = p.WriteUInt32(1).
		WriteUInt32(1).
		WriteInt8(0).
		WriteUInt32(1)

	for i := 0; i < 150; i++ {
		p = p.WriteUInt32(0)
	}

	p = p.WriteUInt32(0).
		WriteUInt32(0).
		WriteUInt32(0)

	return p
}

func MakeAddObj(from *structure.PlayerEntity) net.Packet {
	p := net.StartMergePacket(uint32(from.ID), uint16(net.ADDOBJ), 0xFFFFFF00)
	p = p.WriteUInt8(5).
		WriteUInt32(11).
		WriteUInt8(5).
		WriteUInt32(11).
		WriteUInt16(100).
		WriteFloat32(float32(from.Position.Vec.X)).
		WriteFloat32(float32(from.Position.Vec.Y)).
		WriteFloat32(float32(from.Position.Vec.Z)).
		WriteUInt16(360).
		WriteUInt32(uint32(from.ID)).
		WriteUInt16(0).
		WriteUInt8(1).
		WriteUInt32(230).
		WriteUInt32(0).
		WriteUInt32(0).
		WriteUInt8(1).
		WriteInt32(-1).
		WriteString(from.Name).
		WriteUInt8(from.Gender).
		WriteUInt8(uint8(from.SkinSetID)).
		WriteUInt8(uint8(from.HairID)).
		WriteUInt32(from.HairColor).
		WriteUInt8(uint8(from.FaceID)).
		WriteUInt32(uint32(from.ID)).
		WriteUInt8(1).
		WriteUInt16(0).
		WriteUInt16(0).
		WriteUInt16(0).
		WriteUInt16(0).
		WriteUInt16(uint16(from.Level)).
		WriteInt32(-1).
		WriteUInt32(0).
		WriteUInt8(0).
		WriteUInt32(0).
		WriteUInt8(0).
		WriteUInt8(100).
		WriteUInt32(0).
		WriteUInt32(0x000001F6).
		WriteUInt32(0).
		WriteUInt32(0).
		WriteUInt32(0).
		WriteUInt32(0).
		WriteUInt32(0).
		WriteUInt8(0).
		WriteInt32(-1)

	for i := 0; i < 31; i++ {
		p = p.WriteInt32(-1)
	}
	for i := 0; i < 28; i++ {
		p = p.WriteUInt32(0)
	}

	p = p.WriteUInt8(0).
		WriteInt32(-1).
		WriteUInt32(0)

	return p
}

func MakeDeleteObj(wc *structure.PlayerEntity) net.Packet {
	return net.StartMergePacket(uint32(wc.ID), uint16(0x00F1), 0xFFFFFF00)
}
