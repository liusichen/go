package main

import (
	"encoding/json"
	"fmt"
	"log"
)

//Movie 1
type Movie struct {
	Title string
	Year  int  `json:"relesed"`
	Color bool `json:"color,omitempty"`
	Actor []string
}

var movies = []Movie{
	{Title: "Casablance", Year: 1942, Color: false,
		Actor: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title: "Cool Hand Luke", Year: 1967, Color: true,
		Actor: []string{"Paul Newman"}},
	{Title: "Bullitt", Year: 1968, Color: true,
		Actor: []string{"Steve McQueen", "Jacqueline Bisset"}},
}

func main() {
	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", data)

	data, err = json.MarshalIndent(movies, "", "    ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s\n", err)
	}
	fmt.Printf("%s\n", data)
}
