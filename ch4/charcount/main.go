package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)
	var utflen [utf8.UTFMax + 1]int
	invaild := 0

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invaild++
			continue
		}
		counts[r]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Printf("\nutflen\tcount\n")
	for l, n := range utflen {
		fmt.Printf("%d\t%d\n", l, n)
	}
	if invaild > 0 {
		fmt.Printf("\n%d charactors is invaild\n", invaild)
	}
}
