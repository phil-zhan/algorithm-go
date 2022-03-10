package class01

// 交换两个数
func swap(arr []int, i int, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}

// 交换两个数
func swap2(arr []int, i1 int, i2 int) {

	// 两个数一样，就不要再交换了，不然会导致结果为0
	if arr[i1] == arr[i2] {
		return
	}
	arr[i1] = arr[i1] ^ arr[i2]
	arr[i2] = arr[i1] ^ arr[i2]
	arr[i1] = arr[i1] ^ arr[i2]
}

// PrintArr 打印数组
func PrintArr(arr []int) {
	for i := 0; i < len(arr); i++ {
		println(arr[i])
	}
}
