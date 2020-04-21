package main

import (
	"log"
	"sync"
	"sync/atomic"

	"github.com/rbee3u/golab/utils"
)

const (
	repeat          = 10_000_000_000
	slowConcurrency = 1
	fastConcurrency = 2
)

func main() {
	do("slow", slowConcurrency)
	do("fast", fastConcurrency)
}

func do(name string, concurrency int) {
	defer utils.LogElapsed(name)()
	var (
		sum int64
		wg  sync.WaitGroup
	)
	for c := 0; c < concurrency; c++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			var acc int64
			for r := 0; r < repeat; r += concurrency {
				acc += 2
			}
			atomic.AddInt64(&sum, acc)
		}()
	}
	wg.Wait()
	log.Printf("%s sum is %v", name, atomic.LoadInt64(&sum))
}
