package main

import (
	"algorithm-go/src/class21"
	"fmt"
)

func main() {
	probability1 := class21.BobLiveProbability1(50, 50, 6, 6, 10)
	probability2 := class21.BobLiveProbability2(50, 50, 6, 6, 10)
	fmt.Println(probability1)
	fmt.Println(probability2)

	paper1 := class21.CoinsWaySameValueSamePaper([]int{2, 16, 13, 5, 16, 3, 7}, 15)
	paper2 := class21.CoinsWaySameValueSamePaperDp1([]int{2, 16, 13, 5, 16, 3, 7}, 15)
	paper3 := class21.CoinsWaySameValueSamePaperDp2([]int{2, 16, 13, 5, 16, 3, 7}, 15)
	fmt.Println(paper1)
	fmt.Println(paper2)
	fmt.Println(paper3)

	limit1 := class21.CoinsWayNoLimit([]int{22, 6, 24, 28, 7, 17, 23, 4, 9}, 28)
	limit2 := class21.CoinsWayNoLimitDp1([]int{22, 6, 24, 28, 7, 17, 23, 4, 9}, 28)
	limit3 := class21.CoinsWayNoLimitDp2([]int{22, 6, 24, 28, 7, 17, 23, 4, 9}, 28)
	fmt.Println(limit1)
	fmt.Println(limit2)
	fmt.Println(limit3)

	sum1 := class21.MinPathSum1([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}})
	sum2 := class21.MinPathSum1([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}})
	sum3 := class21.MinPathSum1([][]int{{1, 2, 3}, {4, 5, 6}})
	sum4 := class21.MinPathSumDp([][]int{{1, 2, 3}, {4, 5, 6}})
	sum5 := class21.MinPathSumDp2([][]int{{1, 2, 3}, {4, 5, 6}})
	sum6 := class21.MinPathSumDp2([][]int{{9, 1, 4, 8}})
	fmt.Println(sum1)
	fmt.Println(sum2)
	fmt.Println(sum3)
	fmt.Println(sum4)
	fmt.Println(sum5)
	fmt.Println(sum6)
}
