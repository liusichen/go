package prefix 

import "fmt"
import "os"

func main() {
	s := os.Args[1]
	pre := os.Args[2]
	fmt.Println(HasPrefix(s,pre))
}

func HasPrefix(s,prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}
