package main

import (
	"fmt"
	"go/parser"
	"go/token"
)

func main() {
	fset := token.NewFileSet()
	f, _ := parser.ParseFile(fset, "example.go", src, parser.ParseComments)

	for _, c := range f.Comments {
		fmt.Printf("%s: %q\n", fset.Position(c.Pos()), c.Text())
	}
}

var src = `// Package p provides Add function
// ...
package p

// Add adds two ints.
func add(n, m int) int {
	return n + m
}
`
