package main

import (
	"errors"
	"fmt"
	"log"
)

var ErrNegativeFactorial = errors.New("факториал отрицательных чисел не определён")

func main() {
	result, err := factorial(0)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}

func factorial(x int) (int, error) {
	if x < 0 {
		return 0, ErrNegativeFactorial
	}
	result := 1
	for x > 0 {
		result *= x
		x--
	}
	return result, nil
}
