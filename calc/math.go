package calc

import (
	"fmt"
)

func init() {
	fmt.Println("initializing calc...")
}

// Add adds 2 numbers
func Add(numbers ...int) int {
	sum := 0
	for _, num := range numbers {
		sum = sum + num
	}
	return sum
}
