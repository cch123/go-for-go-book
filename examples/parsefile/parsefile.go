package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	fset := token.NewFileSet()
	f, _ := parser.ParseFile(fset, "example.go", src, parser.Mode(0))

	for _, d := range f.Decls {
		ast.Print(fset, d)
		fmt.Println()
	}
}

var src = `package p
import _ "log"
func add(n, m int) {}
`
