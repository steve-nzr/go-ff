package gametimer

import (
	"flyff/world/service/gamemap"
	"time"
)

type gameTimer struct {
}

var Timer = new(gameTimer)

func (t *gameTimer) Start() {
	prev := time.Now().UnixNano()

	for {
		curr := time.Now().UnixNano()
		gtime := (curr - prev) / int64(time.Millisecond)

		gamemap.Manager.Update(gtime)

		prev = time.Now().UnixNano()
		time.Sleep(1 * time.Millisecond)
	}
}
