package main

import (
	"sync/atomic"
)

type Counter struct {
	V uint64
	_ [120]uint64
}

func (c *Counter) Incr() {
	_ = atomic.AddUint64(&c.V, 1)
}

func (c *Counter) Load() uint64 {
	return atomic.LoadUint64(&c.V)
}
