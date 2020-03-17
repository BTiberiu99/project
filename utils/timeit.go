package utils

import (
	"log"
	"time"
)

func Timeit(Name string, f func()) {
	start := time.Now()
	f()
	elapsed := time.Since(start)
	log.Printf("%s took %s", Name, elapsed)
}
