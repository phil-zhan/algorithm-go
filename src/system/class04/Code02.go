package class04

func GetSmallSum(arr *[]int) int {
	if nil == arr || len(*arr) < 2 {
		return 0
	}

	return processForSmallSum(arr, 0, len(*arr)-1)
}

func processForSmallSum(arr *[]int, left, right int) int {
	if left == right {
		return 0
	}

	var mid = left + ((right - left) >> 1)

	return processForSmallSum(arr, left, mid) + processForSmallSum(arr, mid+1, right) + mergeForSmallSum(arr, left, mid, right)
}

func mergeForSmallSum(arr *[]int, left, mid, right int) int {
	help := make([]int, right-left+1)

	index := 0
	p1 := left
	p2 := mid + 1
	res := 0

	for p1 <= mid && p2 <= right {
		if (*arr)[p1] < (*arr)[p2] {

			// 这里是先取再加加，也就是先计算res，不然会出错
			help[index] = (*arr)[p1]
			res += (right - p2 + 1) * ((*arr)[p1])
			p1++
		} else {
			help[index] = (*arr)[p2]
			p2++
		}
		index++
	}

	for p1 <= mid {
		help[index] = (*arr)[p1]
		p1++
		index++
	}

	for p2 <= right {
		help[index] = (*arr)[p2]
		p2++
		index++
	}

	p3 := 0
	for i := left; i <= right; i++ {
		(*arr)[i] = help[p3]
		p3++
	}

	return res
}
