// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 231.

// Pipeline3 demonstrates a finite 3-stage pipeline
// with range, close, and unidirectional channel types.
package main

import (
	"fmt"
	"time"
)

//!+
func counter(out chan<- int) {
	for x := 0; x < 10; x++ {
		out <- x
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}

func printer(in <-chan int) {
	for v := range in {
		if v == 25 {
			time.Sleep(1 * time.Second)
		}
		fmt.Println(v)
	}
}

func main() {
	naturals := make(chan int, 10)
	squares := make(chan int, 10)

	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)
}

//!-