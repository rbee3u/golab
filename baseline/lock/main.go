package main

import (
	"log"
	"sync"
	"time"
)

const (
	concurrency = 2
	repeat      = 100_000_000
)

func main() {
	start := time.Now()
	var (
		mu  sync.Mutex
		sum int64
		wg  sync.WaitGroup
	)
	for c := 0; c < concurrency; c++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for r := 0; r < repeat; r++ {
				mu.Lock()
				sum += 2
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
	elapsed := time.Since(start)
	log.Printf("sum is %v", sum)
	ops := int64(concurrency * repeat / elapsed.Seconds())
	log.Printf("%v operations per second", ops)
}
