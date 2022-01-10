package watcher

import (
	"fmt"
	"testing"
	"time"
)

func TestSen(t *testing.T) {

	ti := time.NewTicker(time.Second * 1)
	go Watcher(func(args ...interface{}) {
		for {
			select {
			case <-ti.C:
				fmt.Println(time.Now().Second())
				if time.Now().Second()%2 == 0 {
					panic("123312")
				}
			default:
				time.Sleep(time.Millisecond * 50)
			}
		}
	})

	time.Sleep(10 * time.Second)
}
