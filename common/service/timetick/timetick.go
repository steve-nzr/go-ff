package timetick

import (
	"time"
)

// Cancellation cancels the channel
type Cancellation bool

// BeginTick in the given channel
func BeginTick(done <-chan Cancellation, c chan<- int, t time.Duration) {
	prev := time.Now().Nanosecond()

	for {
		curr := time.Now().Nanosecond()
		gametime := (curr - prev) / 1000000
		prev = curr

		c <- gametime

		time.Sleep(t)

		select {
		case <-done:
			return
		default:
			break
		}
	}
}
