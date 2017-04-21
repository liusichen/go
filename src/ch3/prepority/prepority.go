package main

import "fmt"

func main() {
	s := 10
	var u uint8 = 255
	var i int8 = 127
	var a uint32
	test := 255
	a = 3
	str := "hello man"
	str1 := "what"
	fmt.Printf("%d\n",s/5*2)
	fmt.Printf("%d\t %d\n",s&^4, a&^4)
	fmt.Printf("%d\t %d\n",^s,^a)
	fmt.Println(u,u+1,u*u)
	fmt.Println(i,i+1,i*i)
	fmt.Println(test*test%(test+1))
	fmt.Println(str > str1)
}

