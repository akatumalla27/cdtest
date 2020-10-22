package main

import (
	"fmt"
)

func main() {
	fmt.Println(Greet())
}

// Greet Greets GitHub Actions
func Greet() string {
	return "Hello GitHub Actions"
}
