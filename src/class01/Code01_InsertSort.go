package class01

// InsertSort 插入排序
func InsertSort(arr []int) {

	if nil == arr || len(arr) == 0 {
		return
	}

	for i := 0; i < len(arr); i++ {
		for j := i - 1; j >= 0; j-- {
			if arr[j] > arr[j+1] {
				swap(arr, j, j+1)
			}
		}
	}
}

// InsertSort2 插入排序2
func InsertSort2(arr []int) {

	if nil == arr || len(arr) == 0 {
		return
	}

	for i := 1; i < len(arr); i++ {
		for j := i - 1; j >= 0; j-- {
			if arr[j] > arr[j+1] {
				swap2(arr, j, j+1)
			}
		}
	}
}
