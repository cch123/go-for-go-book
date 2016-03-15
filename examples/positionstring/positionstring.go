package main

import (
	"fmt"
	"go/token"
)

func main() {
	fmt.Println("Invalid position without file name:", token.Position{}.String())
	fmt.Println("Invalid position with file name:   ", token.Position{Filename: "example.go"}.String())
	fmt.Println("Valid position without file name:  ", token.Position{Line: 2, Column: 3}.String())
	fmt.Println("Valid position with file name:     ", token.Position{Filename: "example.go", Line: 2, Column: 3}.String())
}
