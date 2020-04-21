package main

import (
	"log"
	"sort"

	"github.com/rbee3u/golab/utils"
)

const (
	repeat = 3_000
	length = 1_000_000
)

func main() {
	array := make([]int, length)
	for i := 0; i < length; i++ {
		array[i] = (i & 3) - 2
	}
	do("slow", array)
	sort.Ints(array)
	do("fast", array)
}

func do(name string, array []int) {
	defer utils.LogElapsed(name)()
	var sum int
	for r := 0; r < repeat; r++ {
		for i := 0; i < length; i++ {
			if array[i] >= 0 {
				sum += array[i]
			}
		}
	}
	log.Printf("%s sum is %v", name, sum)
}
