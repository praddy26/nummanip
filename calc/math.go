package calc

import (
	"fmt"
)

func init() {
	fmt.Println("initializing calc...")
}

// Add adds 2 numbers
func Add(i, j int) int {
	return i + j
}
