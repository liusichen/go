package main

import "fmt"
import "bufio"
import "os"

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanLines)
	for input.Scan() {
		counts[input.Text()]++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Println(n, line)
		}
	}

}
