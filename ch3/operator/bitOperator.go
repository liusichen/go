package main

import "fmt"

func main() {
	var x uint8 = 1<<6 | 1<<4
	var y uint8 = 1<<3 | 1<<4

	fmt.Printf("x:%08b\t y:%08b\n",x,y)
	fmt.Printf("%08b\n",x&y)
	fmt.Printf("%08b\n",x|y)
	fmt.Printf("%08b\n",x^y)
	fmt.Printf("%08b\n",x&^y)

}

