package main

import (
	"fmt"
)

func routine(i chan int) {
	fmt.Printf("routine %v finished\n", <-i)
}

func main() {
	num := make(chan int)
	for i := 0; i < 10; i++ {
		go routine(num)
		num <- i
	}
	fmt.Println("main finished")
}
