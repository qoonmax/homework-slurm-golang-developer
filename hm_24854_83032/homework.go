package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			stack := debug.Stack()
			fmt.Printf("Panic: %v\nStack trace:\n%s\n", r, stack)
		}
	}()

	riskyOperation()
}

func riskyOperation() {
	panic("что-то пошло не так")
}
