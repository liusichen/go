package main

import "fmt"

func main() {
	var x, y int
	fmt.Scanf("%d %d", &x, &y)
	fmt.Println(gcd(x, y))
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}
