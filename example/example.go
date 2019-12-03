package main

import (
	"fmt"

	"github.com/invisiblethreat/locate"
)

func main() {
	fmt.Printf("You're in %s\n", locate.WhereAmI())
	example()
}

func example() {
	fmt.Printf("You're in %s\n", locate.WhereAmI())
}
