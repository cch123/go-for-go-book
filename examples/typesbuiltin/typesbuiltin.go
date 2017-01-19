package main

import (
	"fmt"
	"go/types"
)

func main() {
	obj := types.Universe.Lookup("append")
	fmt.Printf("%v (%T)\n", obj, obj)
	fmt.Printf("%v (%T)\n", obj.Type(), obj.Type())
}
