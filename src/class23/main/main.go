package main

import (
	"algorithm-go/src/class23"
	"fmt"
	"math/rand"
	"time"
)

func main() {

	times := 100
	fmt.Println("测试开始")
	for i := 0; i < times; i++ {
		arr := randomArray(20, 1, 10000)
		closed1 := class23.SplitSumClosed(*arr)
		closed2 := class23.SplitSumClosedDP(*arr)

		if closed1 != closed2 {
			fmt.Println("测试不通过")
			fmt.Println("arr", *arr)
			fmt.Println("closed1", closed1)
			fmt.Println("closed2", closed2)
		}
	}

	for i := 0; i < times; i++ {
		arr := randomArray(6, 1, 10)
		closed3 := class23.SplitSumClosedSizeHalf(*arr)
		closed4 := class23.SplitSumClosedSizeHalfDP(*arr)
		if closed3 != closed4 {
			fmt.Println("测试不通过")
			fmt.Println("arr", *arr)
			fmt.Println("closed3", closed3)
			fmt.Println("closed4", closed4)
		}
	}
	fmt.Println("测试完成")

}

/**
 * 随机数组
 */
func randomArray(len int, min int, max int) *[]int {
	rand.Seed(time.Now().Unix())
	length := rand.Intn(len)

	arr := make([]int, length)

	for i := 0; i < length; i++ {
		arr[i] = rand.Intn(max-min) + min
	}
	return &arr
}
