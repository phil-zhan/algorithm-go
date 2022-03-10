package class04

func BiggerTwice(arr *[]int) int {
	if nil == arr || len(*arr) < 2 {
		return 0
	}

	return processForBiggerTwice(arr, 0, len(*arr)-1)
}

func processForBiggerTwice(arr *[]int, left, right int) int {
	if left == right {
		return 0
	}

	mid := left + ((right - left) >> 1)

	return processForBiggerTwice(arr, left, mid) + processForBiggerTwice(arr, mid+1, right) + mergeForBiggerTwice(arr, left, mid, right)
}

func mergeForBiggerTwice(arr *[]int, left, mid, right int) int {
	windowR := mid + 1
	p4 := left
	ans := 0

	for p4 <= mid {
		for windowR <= right && (*arr)[p4] > ((*arr)[windowR]*2) {
			windowR++
		}
		p4++
		ans += windowR - mid - 1
	}

	help := make([]int, right-left+1)
	index := right - left
	p1 := mid
	p2 := right

	for p1 >= left && p2 >= mid+1 {
		if (*arr)[p2] >= (*arr)[p1] {
			help[index] = (*arr)[p2]
			p2--
		} else {
			help[index] = (*arr)[p1]
			p1--
		}

		index--
	}

	for p1 >= left {
		help[index] = (*arr)[p1]
		index--
		p1--
	}
	for p2 >= mid+1 {
		help[index] = (*arr)[p2]
		p2--
		index--
	}

	p3 := left
	p5 := 0
	for p3 <= right {
		(*arr)[p3] = help[p5]
		p3++
		p5++
	}

	return ans
}
