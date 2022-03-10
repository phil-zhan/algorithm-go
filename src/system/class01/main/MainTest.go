package main

import (
	"algorithm-go/src/system/class01"
)

func main() {

	var arr = []int{1, 2, 0, 4, 6, 1}
	var arr2 = []int{1, 2, 0, 4, 6, 1}
	var arr3 = []int{1, 2, 0, 4, 6, 1}
	var arr4 = []int{1, 2, 0, 4, 6, 1}

	// 插入排序
	class01.InsertSort(arr)
	class01.PrintArr(arr)

	println("================================")

	// 插入排序2
	class01.InsertSort2(arr2)
	class01.PrintArr(arr2)

	println("================================")

	// 选择排序
	class01.SelectedSort(arr3)
	class01.PrintArr(arr3)

	println("================================")

	// 冒泡排序
	class01.BubbleSort(arr4)
	class01.PrintArr(arr4)
}
