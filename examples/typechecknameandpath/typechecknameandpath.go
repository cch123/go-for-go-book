package main

import (
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"path/filepath"
	"runtime"
)

func main() {
	path := "cmd/cover"

	fset := token.NewFileSet()
	aPkgs, _ := parser.ParseDir(fset, filepath.Join(runtime.GOROOT(), "src", path), nil, parser.Mode(0))

	conf := types.Config{Importer: importer.Default()}

	for _, aPkg := range aPkgs {
		files := []*ast.File{}
		for _, f := range aPkg.Files {
			files = append(files, f)
		}
		pkg, _ := conf.Check(path, fset, files, nil)
		fmt.Printf("path=%v name=%v\n", pkg.Path(), pkg.Name())
	}
}
