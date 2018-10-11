package in

import (
	"flyff/core"
	"flyff/world/game/component"
	"flyff/world/game/structure"
	"flyff/world/packets/format"
	"flyff/world/packets/out"
	"flyff/world/service/gamemap"
	"fmt"
)

func Join(wc *structure.WorldClient, p *core.Packet) {
	var join format.Join
	join.Construct(p)

	c := new(core.Character)
	db := core.GetDbConnection()
	db.First(c, join.PlayerID)

	if c == nil {
		fmt.Println("Player not exist !")
		return
	}

	wc.PlayerEntity = new(structure.PlayerEntity)
	wc.PlayerEntity.WorldClient = wc
	wc.PlayerEntity.ID = core.GenerateID()
	wc.PlayerEntity.PlayerID = uint32(c.ID)
	wc.PlayerEntity.Slot = c.Slot
	wc.PlayerEntity.HairColor = c.HairColor
	wc.PlayerEntity.HairID = c.HairID
	wc.PlayerEntity.FaceID = c.FaceID
	wc.PlayerEntity.SkinSetID = c.SkinSetID
	wc.PlayerEntity.JobID = c.JobID
	wc.PlayerEntity.Angle = 360.0
	wc.PlayerEntity.Gender = c.Gender
	wc.PlayerEntity.Level = c.Level
	wc.PlayerEntity.Type = 5 // Mover
	wc.PlayerEntity.Size = 100
	wc.PlayerEntity.Position = component.Position{
		MapID: c.MapID,
		X:     c.PosX,
		Y:     c.PosY,
		Z:     c.PosZ,
	}
	wc.PlayerEntity.Name = c.Name
	if c.Gender == 0 {
		wc.PlayerEntity.ModelID = 11
	} else if c.Gender == 1 {
		wc.PlayerEntity.ModelID = 12
	}
	wc.PlayerEntity.Statistics = component.Statistics{
		Strength:     c.Strength,
		Stamina:      c.Stamina,
		Dexterity:    c.Dexterity,
		Intelligence: c.Intelligence,
	}

	wc.Send(out.MakeSpawn(wc.PlayerEntity))
	gamemap.Manager.Register(wc)
}
