package move

import (
	"flyff/common/service/cache"
	"flyff/common/service/timetick"
	"fmt"
	"time"

	"github.com/golang/geo/r3"
)

// ProcessMove for the given (NetClientID) player
func ProcessMove(id uint32, destAngle r3.Vector) {
	done := make(chan timetick.Cancellation)
	tick := make(chan int)
	go timetick.BeginTick(done, tick, 150*time.Millisecond)

	for {
		t := <-tick
		p := cache.FindByNetID(id)
		if p == nil {
			return
		}

		fmt.Println(t)
	}
}
