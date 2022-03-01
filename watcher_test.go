package watcher

import (
	"math/rand"
	"testing"
	"time"
)

func TestWatcher(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	//	Retry 10 times, the retry interval is 1 second,
	//	after 10 seconds of stable operation,
	//	reset the number of retries
	w := NewWatcher(10, time.Second*1, time.Second*10)
	w.On(func(args ...interface{}) {
		for {
			if rand.Int()%9 == 0 {
				panic("crash")
			}
			time.Sleep(500 * time.Millisecond)
		}
	})
}
