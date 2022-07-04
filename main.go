package main

import (
	"fmt"
	"hui0000/go_algorithm/algorithm"
)

func main() {
	s := algorithm.BuildSegTree(10)
	s.RangeAdd(0, 5, 2)
	s.RangeAdd(3, 6, 1)

	fmt.Println(s.GetRange(3, 6))
}
