package gamemap

import (
	"flyff/world/game/structure"
)

type gameMap structure.GameMap

func (gm *gameMap) Update(time int64) {
	for _, p := range gm.Players {
		for _, us := range gm.UpdatableSystems {
			us.Update(time, p)
		}
	}
}
