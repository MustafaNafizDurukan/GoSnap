package counter

import (
	"time"
)

// Start loops for the specified duration.
// The fn function is called every tick, and the loop is stopped
// when the duration is reached.
func Start(duration time.Duration, fn func()) {
	ticker := time.NewTicker(time.Second / 5)
	timer := time.NewTimer(duration)

	for {
		select {
		case <-ticker.C:
			fn()
		case <-timer.C:
			return
		}
	}
}
