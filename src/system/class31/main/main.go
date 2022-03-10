package main

import (
	sa "algorithm-go/src/system/class31"
	"fmt"
)

func main() {

	// segmentTree test
	var s1 = &sa.SegmentTree{}
	arr := []int{1, 2, 3, 4, 5, 6, 7, 1, 2, 3, 4, 5, 6, 7}
	s1.Init(arr)
	res := s1.Query(1, 4, 1, 14, 1)
	fmt.Println(res)
	s1.Add(2, 3, 1, 1, 14, 1)

	res1 := s1.Query(1, 14, 1, 14, 1)
	fmt.Println(res1)

	s1.Update(1, 5, 1, 1, 14, 1)
	res2 := s1.Query(1, 14, 1, 14, 1)
	fmt.Println(res2)
}
