package main

import (
	"log"

	"github.com/rbee3u/golab/utils"
)

const (
	length   = 499_999_999
	slowStep = 20_000
	fastStep = 1
)

func main() {
	array := make([]int, length)
	for i := 0; i < length; i++ {
		array[i] = i
	}
	do("slow", array, slowStep)
	do("fast", array, fastStep)
}

func do(name string, array []int, step int) {
	defer utils.LogElapsed(name)()
	var sum int64
	for i := step; i != 0; {
		sum += int64(array[i])
		i = (i + step) % length
	}
	log.Printf("%s sum is %v", name, sum)
}
