package main

import (
	"fmt"

	"github.com/invisiblethreat/locate"
)

func main() {
	fmt.Printf("package.function:file:line %s\n", locate.Location())
	example()
}

func example() {
	fmt.Printf("package.function:file:line %s\n", locate.Location())
}
