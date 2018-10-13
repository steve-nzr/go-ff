package movement

import (
	"flyff/world/entities"
	"fmt"
	"math"

	"github.com/golang/geo/r3"
)

type System struct {
}

func (s *System) Update(time int64, pe *entities.PlayerEntity) {
	if pe.Moving.X != 0 {
		if pe.Moving.Distance(pe.Position.Vec) < 0.1 {
			pe.Moving.Vector = r3.Vector{}
			fmt.Println("ARRIVE")
		} else {
			var speed = (0.1 * 100.0) * float64(float64(time)/1000.0)
			distX := pe.Moving.X - pe.Position.Vec.X
			distZ := pe.Moving.Z - pe.Position.Vec.Z
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

			pe.Position.Vec.X += offsetX
			pe.Position.Vec.Z += offsetZ
		}
	}
}
