package main

import (
	"log"

	"github.com/rbee3u/golab/utils"
)

const (
	repeat = 10_000_000_000
	step   = 8
)

func main() {
	slow()
	fast()
}

func slow() {
	defer utils.LogElapsed("slow")()
	var sum int64
	for r := 0; r < repeat; r++ {
		sum += int64(r & 1)
	}
	log.Printf("slow sum is %v", sum)
}

func fast() {
	defer utils.LogElapsed("fast")()
	var counters [step]int64
	for r := 0; r < repeat; r += step {
		counters[0] += int64((r + 0) & 1)
		counters[1] += int64((r + 1) & 1)
		counters[2] += int64((r + 2) & 1)
		counters[3] += int64((r + 3) & 1)
		counters[4] += int64((r + 4) & 1)
		counters[5] += int64((r + 5) & 1)
		counters[6] += int64((r + 6) & 1)
		counters[7] += int64((r + 7) & 1)
	}
	var sum int64
	for slot := 0; slot < step; slot++ {
		sum += counters[slot]
	}
	log.Printf("fast sum is %v", sum)
}
