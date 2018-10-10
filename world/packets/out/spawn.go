package out

import (
	"flyff/core"
	"flyff/world/game/structure"
	"math"
)

func MakeSpawn(from *structure.WorldClient) core.Packet {
	p := core.StartMergePacket(uint32(from.Character.ID), uint16(core.ENVIRONMENTALL), 0x0000FF00).
		WriteUInt32(0)

	p = p.AddMergePart(core.WORLDREADINFO, uint32(from.Character.ID)).
		WriteUInt32(from.Character.MapID).
		WriteFloat32(from.Character.PosX).
		WriteFloat32(from.Character.PosY).
		WriteFloat32(from.Character.PosZ)

	p = p.AddMergePart(core.ADDOBJ, uint32(from.Character.ID))

	p = p.WriteUInt8(5)
	if from.Character.Gender == 0 {
		p = p.WriteUInt32(11)
	} else if from.Character.Gender == 1 {
		p = p.WriteUInt32(12)
	}

	p = p.WriteUInt8(5)
	if from.Character.Gender == 0 {
		p = p.WriteUInt32(11)
	} else if from.Character.Gender == 1 {
		p = p.WriteUInt32(12)
	}

	p = p.WriteUInt16(100).
		WriteFloat32(from.Character.PosX).
		WriteFloat32(from.Character.PosY).
		WriteFloat32(from.Character.PosZ).
		WriteUInt16(360).
		WriteUInt32(uint32(from.Character.ID))

	p = p.WriteUInt16(0).
		WriteUInt8(1).
		WriteUInt32(230).
		WriteUInt32(0).
		WriteUInt32(0).
		WriteUInt8(1).
		WriteInt32(-1)

	p = p.WriteString(from.Character.Name).
		WriteUInt8(from.Character.Gender).
		WriteUInt8(uint8(from.Character.SkinSetID)).
		WriteUInt8(uint8(from.Character.HairID)).
		WriteUInt32(from.Character.HairColor).
		WriteUInt8(uint8(from.Character.FaceID)).
		WriteUInt32(uint32(from.Character.ID)). // Playerdata ID
		WriteUInt8(uint8(from.Character.JobID)).
		WriteUInt16(from.Character.Strength).
		WriteUInt16(from.Character.Stamina).
		WriteUInt16(from.Character.Dexterity).
		WriteUInt16(from.Character.Intelligence).
		WriteUInt16(from.Character.Level)

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
	p = p.WriteUInt32(from.Character.MapID).
		WriteFloat32(from.Character.PosX).
		WriteFloat32(from.Character.PosY).
		WriteFloat32(from.Character.PosZ)

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

	p = p.WriteUInt8(from.Character.Slot)
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

func MakeAddObj(from *structure.WorldClient) core.Packet {
	p := core.StartMergePacket(uint32(from.Character.ID), uint16(core.ADDOBJ), 0xFFFFFF00)
	p = p.WriteUInt8(5).
		WriteUInt32(11).
		WriteUInt8(5).
		WriteUInt32(11).
		WriteUInt16(100).
		WriteFloat32(from.Character.PosX).
		WriteFloat32(from.Character.PosY).
		WriteFloat32(from.Character.PosZ).
		WriteUInt16(360).
		WriteUInt32(uint32(from.Character.ID)).
		WriteUInt16(0).
		WriteUInt8(1).
		WriteUInt32(230).
		WriteUInt32(0).
		WriteUInt32(0).
		WriteUInt8(1).
		WriteInt32(-1).
		WriteString(from.Character.Name).
		WriteUInt8(from.Character.Gender).
		WriteUInt8(uint8(from.Character.SkinSetID)).
		WriteUInt8(uint8(from.Character.HairID)).
		WriteUInt32(from.Character.HairColor).
		WriteUInt8(uint8(from.Character.FaceID)).
		WriteUInt32(uint32(from.Character.ID)).
		WriteUInt8(1).
		WriteUInt16(0).
		WriteUInt16(0).
		WriteUInt16(0).
		WriteUInt16(0).
		WriteUInt16(from.Character.Level).
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

func MakeDeleteObj(wc *structure.WorldClient) core.Packet {
	return core.StartMergePacket(uint32(wc.Character.ID), uint16(0x00F1), 0xFFFFFF00)
}
