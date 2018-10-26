package move

import (
	"flyff/common/service/cache"
	"flyff/common/service/timetick"
	"fmt"
	"math"
	"time"

	"github.com/golang/geo/r3"
)

func movePlayerByDest(p *cache.Player, t int) {
	if t < 0 {
		return
	}

	if p.Moving.Destination.Distance(p.Position.Vec) < 0.1 {
		p.Moving.Destination = r3.Vector{}
		p.Moving.Motion = 0
		fmt.Println("Arrive")
		SaveMovingComponent(p)
		return
	}

	var speed = (0.08 * 100.0) * (float64(t) / 1000.0)
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

	SavePosition(p)
}

// ProcessDestPosMove for the given (NetClientID) player
func ProcessDestPosMove(id uint32, destPos r3.Vector) {
	done := make(chan timetick.Cancellation)
	tick := make(chan int)
	go timetick.BeginTick(done, tick, 150*time.Millisecond)

	for {
		t := <-tick
		p := cache.FindByNetID(id)
		if p == nil {
			return
		}

		if p.Moving.Destination != destPos {
			done <- true
			return
		}

		movePlayerByDest(p, t)
	}
}
