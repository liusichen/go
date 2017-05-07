package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n)
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x == 0 || x == 1 {
		return 0
	}
	var fib = make([]int, x+1)
	fib[0] = 0
	fib[1] = 1
	for n := 2; n <= x; n++ {
		fib[n] = fib[n-1] + fib[n-2]
	}
	return fib[x]
}
