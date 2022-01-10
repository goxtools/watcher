package watcher

import (
	"github.com/goxtools/watcher/log"
	"go.uber.org/zap"
)

func Watcher(f func(args ...interface{})) {
	reWatcher(f)
}

func reWatcher(f func(args ...interface{})) {
	defer func() {
		if e := recover(); e != nil {
			log.Error("[watch] Exception detected", zap.Any("e", e))
			Watcher(f)
		}
	}()

	f()
}
