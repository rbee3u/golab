package utils

import (
	"log"
	"time"
)

func LogElapsed(name string) func() {
	start := time.Now()
	return func() {
		elapsed := time.Since(start)
		log.Printf("%s completed in %s", name, elapsed)
	}
}
