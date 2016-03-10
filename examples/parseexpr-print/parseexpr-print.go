package main

import (
	"go/ast"
	"go/parser"
)

func main() {
	expr, _ := parser.ParseExpr("a * -1")
	ast.Print(nil, expr)
}
