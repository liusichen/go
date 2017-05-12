package main

import "fmt"

func sum(val ...int) int {
	total := 0
	for _, val := range val {
		total += val
	}
	return total
}

func main() {
	fmt.Println(sum())
	fmt.Println(sum(2))
	fmt.Println(sum(4))
	fmt.Println(sum(1, 2, 3, 4))

	values := []int{1, 2, 3, 4}
	fmt.Println(sum(values...))
}
