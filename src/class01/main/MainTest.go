package main

import sorter "algorithm-go/src/class01"

func main() {

	var arr = []int{1, 2, 0, 4, 6, 1}
	var arr2 = []int{1, 2, 0, 4, 6, 1}
	var arr3 = []int{1, 2, 0, 4, 6, 1}
	var arr4 = []int{1, 2, 0, 4, 6, 1}

	// 插入排序
	sorter.InsertSort(arr)
	sorter.PrintArr(arr)

	println("================================")

	// 插入排序2
	sorter.InsertSort2(arr2)
	sorter.PrintArr(arr2)

	println("================================")

	// 选择排序
	sorter.SelectedSort(arr3)
	sorter.PrintArr(arr3)

	println("================================")

	// 冒泡排序
	sorter.BubbleSort(arr4)
	sorter.PrintArr(arr4)
}
