package main

import (
	"log"
	"math/rand"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Fatal("Panic")
		}
	}()

	_ = 10 / (rand.Int() * 0)
}
