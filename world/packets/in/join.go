package in

import (
	"flyff/core"
	"flyff/core/net"
	"flyff/world/entities"
	movementComponent "flyff/world/feature/movement/component"
	"flyff/world/game/component"
	"flyff/world/packets/format"
	"flyff/world/packets/out"
	"flyff/world/service/gamemap"
	"fmt"

	"github.com/golang/geo/r3"
)

func Join(pe *entities.PlayerEntity, p *net.Packet) {
	var join format.Join
	join.Construct(p)

	c := new(core.Character)
	db := core.GetDbConnection()
	db.First(c, join.PlayerID)

	if c == nil {
		fmt.Println("Player not exist !")
		return
	}

	pe.ID = net.GenerateID()
	pe.PlayerID = uint32(c.ID)
	pe.Slot = c.Slot
	pe.HairColor = c.HairColor
	pe.HairID = c.HairID
	pe.FaceID = c.FaceID
	pe.SkinSetID = c.SkinSetID
	pe.JobID = c.JobID
	pe.Angle = 360.0
	pe.Gender = c.Gender
	pe.Level = c.Level
	pe.Type = 5 // Mover
	pe.Size = 100
	pe.Position = movementComponent.Position{
		MapID: c.MapID,
		Vec: r3.Vector{
			X: float64(c.PosX),
			Y: float64(c.PosY),
			Z: float64(c.PosZ),
		},
	}
	pe.Name = c.Name
	if c.Gender == 0 {
		pe.ModelID = 11
	} else if c.Gender == 1 {
		pe.ModelID = 12
	}

	pe.Statistics = component.Statistics{
		Strength:     c.Strength,
		Stamina:      c.Stamina,
		Dexterity:    c.Dexterity,
		Intelligence: c.Intelligence,
	}

	pe.Send(out.MakeSpawn(pe))
	gamemap.Manager.Register(pe)
}
