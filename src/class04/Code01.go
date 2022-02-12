package class04

func MergeSort(arr *[]int) {
	if nil == arr || len(*arr) == 0 {
		return
	}
	process(arr, 0, len(*arr)-1)
}

func process(arr *[]int, left int, right int) {
	if left == right {
		return
	}

	var mid = left + ((right - left) >> 1)
	process(arr, left, mid)
	process(arr, mid+1, right)

	merge(arr, left, mid, right)
}

func merge(arr *[]int, left int, mid int, right int) {
	help := make([]int, right-left+1)

	index := 0
	p1 := left
	p2 := mid + 1

	for p1 <= mid && p2 <= right {
		if (*arr)[p1] <= (*arr)[p2] {
			help[index] = (*arr)[p1]
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
}
