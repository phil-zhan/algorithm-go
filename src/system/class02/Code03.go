package class02

// GetOddTimesNum 在一个数组中，只有两个数出现了奇数次，其他都出现了偶数次
func GetOddTimesNum(arr []int) []int {

	var eor = 0

	for i := 0; i < len(arr); i++ {
		eor ^= arr[i]
	}

	// 使用最右侧的一个1【在这两个基数中，肯定有一个基数的该位置是1，另外一个基数的该位置是0】
	var rightOnePosNum = getRightNonZeroPoseNum(eor)

	// 过滤一部分异或运算
	var oneOnly = 0
	for i := 0; i < len(arr); i++ {
		if (arr[i] & rightOnePosNum) != 0 {
			oneOnly ^= arr[i]
		}
	}
	return []int{oneOnly, oneOnly ^ eor}
}

// GetOddTimesNum2 在一个数组中，只有两个数出现了奇数次，其他都出现了偶数次
func GetOddTimesNum2(arr []int) []int {
	var eor = 0
	for i := 0; i < len(arr); i++ {
		eor ^= arr[i]
	}

	rightPosNum := getRightNonZeroPoseNum(eor)
	var onlyOne = 0
	for i := 0; i < len(arr); i++ {
		if arr[i]&rightPosNum != 0 {
			onlyOne ^= arr[i]
		}
	}

	return []int{onlyOne, onlyOne ^ eor}
}

// getRightNonZeroPoseNum 获取最右侧的1
func getRightNonZeroPoseNum(i int) int {

	// 在 go 中不能使用 ~ 取反
	// return i & (~i + 1)
	// return i & (^i + 1)
	return i & (-i)
}
