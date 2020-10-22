package main

import (
	"fmt"
	"testing"
)

func main() {
	fmt.Println(Greet())
}

// Greet Greets GitHub Actions
func Greet() string {
	return "Hello GitHub Actions"
}
func TestGreetsGitHub(t *testing.T) {
	result := Greet()
	if result != "Hello GitHub Actions" {
		t.Errorf("Greet() = %s; want Hello GitHub actions", result)
	}
}
