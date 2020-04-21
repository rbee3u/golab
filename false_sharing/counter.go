package main

import (
	"sync/atomic"
)

type Counter interface {
	IncrA()
	IncrB()
}

type SlowCounter struct {
	A uint64
	B uint64
}

func (c *SlowCounter) IncrA() {
	_ = atomic.AddUint64(&c.A, 1)
}

func (c *SlowCounter) IncrB() {
	_ = atomic.AddUint64(&c.B, 1)
}

type FastCounter struct {
	A uint64
	_ [120]uint64
	B uint64
	_ [120]uint64
}

func (c *FastCounter) IncrA() {
	_ = atomic.AddUint64(&c.A, 1)
}

func (c *FastCounter) IncrB() {
	_ = atomic.AddUint64(&c.B, 1)
}
