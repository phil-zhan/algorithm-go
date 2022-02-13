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
}
