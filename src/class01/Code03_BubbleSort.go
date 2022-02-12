package class01

// BubbleSort 冒泡排序
func BubbleSort(arr []int) {

	if nil == arr || len(arr) == 0 {
		return
	}

	for i := len(arr) - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				swap2(arr, j, j+1)
			}
		}
	}
}
