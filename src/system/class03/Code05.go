package class03

import "math"

// 递归获取数组最大值 //

func GetMax(arr []int) int {
	if nil == arr || len(arr) == 0 {
		return -1
	}

	return process(arr, 0, len(arr)-1)
}

func process(arr []int, left int, right int) int {

	if left == right {
		return arr[left]
	}

	var mid = left + ((right - left) >> 1)
	var leftMax = process(arr, left, mid)
	var rightMax = process(arr, mid+1, right)

	return int(math.Max(float64(leftMax), float64(rightMax)))
}
