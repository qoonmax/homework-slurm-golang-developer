package main

import (
	"log"
	"os"
	"runtime/pprof"
)

func main() {
	f, err := os.Create("cpu_panic.prof")
	if err != nil {
		log.Fatal(err)
	}

	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	for i := 0; i < 100000000; i++ {
		division(i, 0)
	}
}

func division(divisible int, divider int) int {
	defer func() {
		recover()
	}()

	result := divisible / divider

	return result
}
