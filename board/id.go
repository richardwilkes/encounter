package board

import "sync/atomic"

var nextID int64 = 1

func NextID() int {
	return int(atomic.AddInt64(&nextID, 1))
}
