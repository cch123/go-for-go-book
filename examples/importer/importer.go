package main

import (
	"fmt"
	"go/importer"
)

func main() {
	pkg, _ := importer.Default().Import("log")
	fmt.Println(pkg.Scope().Names())
	fmt.Println(pkg.Imports())
}
