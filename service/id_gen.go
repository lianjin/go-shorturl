package service

import "sync/atomic"

var (
	counter = int64(0)
)

func GenId() int64 {
	atomic.AddInt64(&counter, int64(1))
	return counter
}
