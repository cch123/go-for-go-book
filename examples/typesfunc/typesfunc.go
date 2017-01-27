package main

import (
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
)

func main() {
	fset := token.NewFileSet()
	f, _ := parser.ParseFile(fset, "example.go", src, parser.Mode(0))

	conf := types.Config{Importer: importer.Default()}

	pkg, _ := conf.Check("path/to/pkg", fset, []*ast.File{f}, nil)

	// type assersions for clarity
	var (
		objF types.Object = pkg.Scope().Lookup("F").(*types.Func)
		objT types.Object = pkg.Scope().Lookup("T").(*types.TypeName)
		objI types.Object = pkg.Scope().Lookup("I").(*types.TypeName)
	) // <1>

	fmt.Println(objF)
	fmt.Println(objT.Type().(*types.Named).Method(0))                  // <2>
	fmt.Println(objI.Type().Underlying().(*types.Interface).Method(0)) // <3>
}

var src = `package p

func F() {}

type T struct{}

func (*T) F() {}

type I interface {
	F()
}
`
