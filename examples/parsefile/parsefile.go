package main

import (
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	var fset *token.FileSet = token.NewFileSet()
	f, _ := parser.ParseFile(fset, "examples/parsefile/parsefile.go", nil, parser.Mode(0))

	for _, d := range f.Decls {
		ast.Print(nil, d)
	}
}
