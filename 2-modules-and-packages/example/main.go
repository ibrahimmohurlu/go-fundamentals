package main

import (
	"fmt"
	"modules-and-packages-example/quote"
)

func main() {
	fmt.Println("Hello from main package in main.go")
	fmt.Println("Random quote:")
	fmt.Println(quote.RandomQuote())
}
