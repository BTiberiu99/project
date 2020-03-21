package utils

import (
	"log"
	"time"
)

func Timeit(Name string, f func()) time.Duration {
	start := time.Now()
	f()
	elapsed := time.Since(start)
	log.Printf("%s took %s", Name, elapsed)
	return elapsed
}
