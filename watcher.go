package watcher

import (
	"github.com/goxtools/watcher/log"
	"go.uber.org/zap"
	"sync/atomic"
	"time"
)

// Watcher Listen for crashes, restart the service
type Watcher struct {
	retryTimes       int32         // expected number of retries
	currentRetryTime int32         // retries remaining
	tf               int32         // the timer has been marked as occupied
	t                *time.Timer   // timer
	retryDelayTime   time.Duration // retry delay time
	resetTime        time.Duration // reset retryTimes, when normal operation for X time,
	// the number of retries returns to the initial value
}

func NewWatcher(retryTimes int32, delayTime time.Duration, resetTime time.Duration) *Watcher {
	return &Watcher{
		retryTimes:       retryTimes,
		currentRetryTime: retryTimes,
		retryDelayTime:   delayTime,
		resetTime:        resetTime,
		tf:               0,
	}
}
func (w *Watcher) On(f func(args ...interface{})) {
	w.on(f)
}
func (w *Watcher) on(f func(args ...interface{})) {
	defer func() {
		if e := recover(); e != nil {
			log.Debug("[Watcher] Watcher Anomaly detected",
				zap.String("sleep", w.retryDelayTime.String()),
				zap.Any("e", e),
			)
			atomic.AddInt32(&w.currentRetryTime, -1)
			if atomic.LoadInt32(&w.currentRetryTime) < 0 {
				log.Debug("The number of retries is gone, stop retrying")
				return
			}
			go func() {
				log.Debug("boom……………………")
				// reset resetTime when a crash
				if atomic.LoadInt32(&w.tf) == 1 && w.t != nil {
					log.Debug("boom, the timer time is reset~")
					w.t.Reset(w.resetTime)
				}
				// already occupied
				if !atomic.CompareAndSwapInt32(&w.tf, 0, 1) {
					log.Debug("The timer has been preempted")
					return
				}
				log.Debug("The timer preempted successfully")
				w.t = time.NewTimer(w.resetTime)
				log.Debug("Initialize the timer",
					zap.String("resetTime", w.resetTime.String()),
				)
				for {
					select {
					case <-w.t.C:
						// Stable operation, reset when time is up
						atomic.StoreInt32(&w.currentRetryTime, w.retryTimes)
						// release
						atomic.CompareAndSwapInt32(&w.tf, 1, 0)
						w.t = nil
						log.Debug("reset timer done")
						log.Debug("reset retries done")
						return
					default:
						log.Debug("default: timer has no data")
						time.Sleep(1 * time.Second)
					}
				}
			}()
			time.Sleep(w.retryDelayTime)
			w.On(f)
		}
	}()

	f()
}
