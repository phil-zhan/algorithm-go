package class01

// SelectedSort ιζ©ζεΊ
func SelectedSort(arr []int) {

	if nil == arr || len(arr) == 0 {
		return
	}

	for i := 0; i < len(arr); i++ {
		var min = i
		for j := i; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		swap2(arr, i, min)
	}

}
