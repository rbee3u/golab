package main

import (
	"log"
	"time"
)

const repeat = 10_000_000_000

func main() {
	start := time.Now()
	var sum int64
	for r := 0; r < repeat; r++ {
		sum += 2
	}
	elapsed := time.Since(start)
	log.Printf("sum is %v", sum)
	ops := int64(repeat / elapsed.Seconds())
	log.Printf("%v operations per second", ops)
}
