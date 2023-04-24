package counter

import (
	"sync"
	"time"
)

// Start loops for the specified duration.
// The fn function is called every tick, and the loop is stopped
// when the duration is reached.
func Start(duration time.Duration, fn func()) {
	ticker := time.NewTicker(time.Second)
	timer := time.NewTimer(duration)

	var wg sync.WaitGroup

	for {
		select {
		case <-ticker.C:
			wg.Add(1)
			go func() {
				defer wg.Done()
				fn()
			}()
		case <-timer.C:
			wg.Wait()
			return
		}
	}
}
