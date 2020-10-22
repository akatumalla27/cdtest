package main

import (
	"fmt"
	"github.com/brpaz/go-github-actions/hello"
)

func main() {
	fmt.Println(Greet())
}

// Greet Greets GitHub Actions
func Greet() string {
	return "Hello GitHub Actions"
}
