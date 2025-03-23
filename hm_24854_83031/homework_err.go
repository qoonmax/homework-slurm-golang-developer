package main

import (
	"errors"
	"log"
	"os"
	"runtime/pprof"
)

var dividerIsZeroError = errors.New("divider is zero")

func main() {
	f, err := os.Create("cpu_err.prof")
	if err != nil {
		log.Fatal(err)
	}

	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	for i := 0; i < 100000000; i++ {
		_, err = division(i, 0)

		if err != nil {
			continue
		}
	}
}

func division(divisible int, divider int) (int, error) {
	if divider == 0 {
		return 0, dividerIsZeroError
	}

	result := divisible / divider

	return result, nil
}
