package main

import (
	class042 "algorithm-go/src/system/class04"
	"fmt"
)

func main() {
	var arr = []int{1, 2, 3, 2, 2, 2, 2, 1, 1, 1, 5, 6, 7, 3}
	class042.MergeSort(&arr)
	fmt.Println("排序后的结果:")
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i])
	}
	fmt.Println()

	fmt.Println("==========================================")
	var arr2 = []int{1, 3, 5, 7, 9, 2, 4, 6, 8, 10}
	sum := class042.GetSmallSum(&arr2)
	fmt.Println("最小和的结果")
	fmt.Println(sum)

	fmt.Println("==========================================")
	pairs := class042.ReversePairs([]int{1, 3, 2, 5, 6, 8})
	fmt.Println("逆序对的结果")
	fmt.Println(pairs)

	fmt.Println("==========================================")
	//arr3 := []int{-10, -30, -50, 17, 8, 52, -64, 0, -17, -40, -12, 0, -17, -30, 33, -18, -6, 61, -52}
	arr3 := []int{2, 4, -8, -13, 24, 53, -26, 27, 7, -73, -88, 59, -77, -24, -1, 1, -8, -4, 30, -27, -52, -13, 17, -53, -31, 30, -42, 26, 16, 16, -58, -25, -40, -26, -40, -82, 19, 25, 40, 80, 41, 88, 48, -52, -38, -14, -20, -45, 90, -13, -2, -42, 14, -5, 9}
	biggerTwice := class042.BiggerTwice(&arr3)
	fmt.Println("大于两倍的结果")
	fmt.Println(biggerTwice)

}
