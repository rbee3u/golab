package main

import (
	"log"
	"time"
)

const (
	repeat = 10_000
	length = 1_000_000
)

func main() {
	array := make([]int64, length)
	for i := 0; i < length; i++ {
		array[i] = 2
	}
	start := time.Now()
	var sum int64
	for r := 0; r < repeat; r++ {
		for i := 0; i < length; i++ {
			sum += array[i]
		}
	}
	elapsed := time.Since(start)
	log.Printf("sum is %v", sum)
	ops := int64(repeat * length / elapsed.Seconds())
	log.Printf("%v operations per second", ops)
}
