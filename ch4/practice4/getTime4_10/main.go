package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"learn/go/ch4/github"
)

var day = time.Second * 3600 * 24
var month = day * 30

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d issues\n", result.TotalCount)
	for _, item := range result.Items {
		if day > time.Since(item.CreatedAt) {
			fmt.Println(time.Since(item.CreatedAt))
			fmt.Printf("%T\t%T\n", month, time.Since(item.CreatedAt))
		}
	}
}
