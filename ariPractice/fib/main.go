package main

import "fmt"

func main() {
	var order int
	fmt.Scanf("%d", &order)
	fmt.Println(fib(order))
}

func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}
