package main

import (
	"sync"

	"github.com/rbee3u/golab/utils"
)

const (
	concurrency = 2
	repeat      = 50_000_000
)

func main() {
	do("slow", &SlowCounter{})
	do("fast", &FastCounter{})
}

func do(name string, counter Counter) {
	defer utils.LogElapsed(name)()
	fnList := []func(Counter){
		Counter.IncrA,
		Counter.IncrB,
	}
	var wg sync.WaitGroup
	for i := range fnList {
		fn := fnList[i]
		for c := 0; c < concurrency; c++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for r := 0; r < repeat; r++ {
					fn(counter)
				}
			}()
		}
	}
	wg.Wait()
}
