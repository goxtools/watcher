package internal

import (
	"runtime"
)

func Stack() string {
	stack := make([]byte, 1024*1024)
	n := runtime.Stack(stack, true)
	return string(stack[:n])
}
