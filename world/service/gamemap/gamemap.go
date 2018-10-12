package gamemap

import (
	"flyff/world/game/structure"
	"fmt"
	"math"

	"github.com/golang/geo/r3"
)

type gameMap structure.GameMap

func (gm *gameMap) Update(time int64) {
	for _, p := range gm.Players {
		// MOVE
		if p.Moving.Destination.X != 0 {
			if p.Moving.Destination.Distance(p.Position.Vec) < 0.1 {
				p.Moving.Destination = r3.Vector{}
				fmt.Println("ARRIVE")
			} else {
				var speed = (0.1 * 100.0) * float64(float64(time)/1000.0)
				distX := p.Moving.Destination.X - p.Position.Vec.X
				distZ := p.Moving.Destination.Z - p.Position.Vec.Z
				distance := math.Sqrt(distX*distX + distZ*distZ)

				deltaX := distX / distance
				deltaZ := distZ / distance
				offsetX := deltaX * speed
				offsetZ := deltaZ * speed

				if math.Abs(offsetX) > math.Abs(distX) {
					offsetX = distX
				}
				if math.Abs(offsetZ) > math.Abs(distZ) {
					offsetZ = distZ
				}

				p.Position.Vec.X += offsetX
				p.Position.Vec.Z += offsetZ
			}
		}
	}
}
