package in

import (
	"flyff/core"
	"flyff/world/game/structure"
	"flyff/world/packets/format"
	"flyff/world/packets/out"
	"flyff/world/service/map"
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

	wc.Character = c

	wc.Send(out.MakeSpawn(wc))
	mapmanager.Manager.Register(wc)
}
