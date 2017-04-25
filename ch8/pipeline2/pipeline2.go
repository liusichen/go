// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 229.

// Pipeline2 demonstrates a finite 3-stage pipeline.
package main

import (
	"fmt"
	"time"
)

//!+
func main() {
	naturals := make(chan int, 10)
	squares := make(chan int, 8)

	// Counter
	go func() {
		for x := 0; x < 10; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	// Squarer
	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	// Printer (in main goroutine)
	for x := range squares {
		if x == 5 {
			time.Sleep(1 * time.Second)
		}
		fmt.Println(x)
	}
}

//!-
