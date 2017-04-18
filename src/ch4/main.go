package main

import "fmt"

import "ch4/reverse"

func main() {
	s := []int{0, 1, 2, 3, 4, 5, 6}

	fmt.Println(s)
	x := s[2:]
	x = append(x, 12)
	fmt.Println(x)
	reverse.ReverseInt(x)
	s = append(s, 17)
	fmt.Println(x)
	fmt.Println(s)
	reverse.ReverseInt(s[:])
	fmt.Println(s)
}
