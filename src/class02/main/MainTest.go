package main

import (
	"algorithm-go/src/class02"
	"fmt"
)

func main() {
	// 不使用额外变量，交换两个数
	a := 3
	b := 2
	fmt.Println("交换前：", a, b)
	class02.Swap(&a, &b)
	fmt.Println("交换后：", a, b)
	fmt.Println("===============================")

	// 获取一个数的最右侧的一个1
	num := class02.GetRightNonZeroPoseNum(3)
	fmt.Println(num)
	fmt.Println("===============================")

	// 获取出现奇数次的数
	nums1 := class02.GetOddTimesNum([]int{1, 1, 2, 2, 3, 3, 4, 5})
	for i := 0; i < len(nums1); i++ {
		fmt.Println(nums1[i])
	}

	fmt.Println("===============================")
	nums2 := class02.GetOddTimesNum2([]int{1, 1, 2, 2, 3, 3, 4, 5})
	for i := 0; i < len(nums2); i++ {
		fmt.Println(nums2[i])
	}

	// 获取仅出现K次的数
	fmt.Println("===============================")
	arr := []int{6, 6, 2, 2, 2, 2, 3, 3, 3, 3, 5, 5, 5, 5}
	timesNum := class02.GetOnlyKTimesNum(arr, 2, 4)

	fmt.Println(timesNum)
}
