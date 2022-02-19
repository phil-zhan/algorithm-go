package main

import (
	"algorithm-go/src/class48"
	"fmt"
)

func main() {
	largest := class48.FindKthLargest([]int{3, 2, 1, 5, 6, 4}, 2)

	fmt.Println(largest)
}
