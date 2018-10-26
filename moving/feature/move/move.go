package move

import (
	"flyff/common/service/cache"
	"flyff/common/service/timetick"
	"math"
	"time"
)

func movePlayer(p *cache.Player, t int, angle float64) {
	if t < 0 {
		return
	}

	theta := angle * math.Pi / 180

	var distance = (0.08 * 100.0) * (float64(t) / 1000.0)

	p.Position.Vec.X += (math.Sin(theta) * distance)
	p.Position.Vec.Z -= (math.Cos(theta) * distance)

	SavePosition(p)
}

// ProcessMove for the given (NetClientID) player
func ProcessMove(id uint32, angle float64) {
	done := make(chan timetick.Cancellation)
	tick := make(chan int)
	go timetick.BeginTick(done, tick, 150*time.Millisecond)

	for {
		t := <-tick
		p := cache.FindByNetID(id)
		if p == nil {
			return
		}

		if p.Moving.Motion != 5 || p.Moving.Angle != angle {
			return
		}

		movePlayer(p, t, angle)
	}
}
