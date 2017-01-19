package main

import (
	"fmt"
	"go/types"
)

func main() {
	fmt.Println(types.Universe.Names())
}
