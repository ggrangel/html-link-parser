package main

import (
	"fmt"
	"os"

	"github.com/ggrangel/html-link-parser/link"
)

func main() {
	examples := []string{"ex1.html", "ex2.html", "ex3.html", "ex4.html"}

	for _, ex := range examples {
		fmt.Println("Parsing", ex)
		file, err := os.Open(ex)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		links, err := link.Parse(file)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v\n", links)
	}
}
