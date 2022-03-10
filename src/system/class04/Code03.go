package class04

func ReversePairs(nums []int) int {
	if nil == nums || len(nums) < 2 {
		return 0
	}

	return reversePairsProcess(&nums, 0, len(nums)-1)
}

func reversePairsProcess(arr *[]int, left, right int) int {
	if left == right {
		return 0
	}

	mid := left + ((right - left) >> 1)

	return reversePairsProcess(arr, left, mid) + reversePairsProcess(arr, mid+1, right) + reverseMerge(arr, left, mid, right)
}

func reverseMerge(arr *[]int, left, mid, right int) int {
	help := make([]int, right-left+1)

	index := right - left
	p1 := mid
	p2 := right
	curNum := 0

	for p1 >= left && p2 >= mid+1 {
		if (*arr)[p2] >= (*arr)[p1] {
			help[index] = (*arr)[p2]
			p2--
		} else {
			help[index] = (*arr)[p1]
			curNum += p2 - mid
			p1--
		}

		index--
	}
	for p1 >= left {
		help[index] = (*arr)[p1]
		//curNum++
		p1--
		index--
	}

	for p2 >= mid+1 {
		help[index] = (*arr)[p2]
		p2--
		index--
	}

	p3 := 0
	for i := left; i <= right; i++ {
		(*arr)[i] = help[p3]
		p3++
	}

	return curNum
}
