package class02

// GetRightNonZeroPoseNum 获取一个数最右侧的1
func GetRightNonZeroPoseNum(i int) int {

	// 在 go 中不能使用 ~ 取反
	// return i & (~i + 1)
	// return i & (^i + 1)
	return i & (-i)
}
