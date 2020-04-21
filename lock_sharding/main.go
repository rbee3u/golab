package main

import (
	"log"
	"sync"

	"github.com/rbee3u/golab/utils"
)

const (
	concurrency = 8
	repeat      = 50_000_000
	slowShard   = 1
	fastShard   = 8
)

func main() {
	do("slow", slowShard)
	do("fast", fastShard)
}

func do(name string, shard int) {
	defer utils.LogElapsed(name)()
	counters := make([]Counter, shard)
	var wg sync.WaitGroup
	for c := 0; c < concurrency; c++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for r := 0; r < repeat; r++ {
				slot := r % shard
				counters[slot].Incr()
			}
		}()
	}
	wg.Wait()

	var sum int64
	for slot := 0; slot < shard; slot++ {
		sum += int64(counters[slot].Load())
	}
	log.Printf("%s sum is %v", name, sum)
}
