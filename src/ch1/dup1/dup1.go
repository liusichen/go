package main

import "fmt"
import "bufio"
import "os"

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	var ceil int
	for input.Scan()&&ceil < 10  {
		counts[input.Text()]++
		ceil++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Println(n, line)
		}
	}

}
