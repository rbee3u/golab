package main

import (
	"log"
	"sync"
	"sync/atomic"
	"time"
)

const (
	concurrency = 2
	repeat      = 100_000_000
)

func main() {
	start := time.Now()
	var (
		sum int64
		wg  sync.WaitGroup
	)
	for c := 0; c < concurrency; c++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for r := 0; r < repeat; r++ {
				atomic.AddInt64(&sum, 2)
			}
		}()
	}
	wg.Wait()
	elapsed := time.Since(start)
	log.Printf("sum is %v", atomic.LoadInt64(&sum))
	ops := int64(concurrency * repeat / elapsed.Seconds())
	log.Printf("%v operation per second", ops)
}
