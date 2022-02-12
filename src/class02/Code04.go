package class02

// GetOnlyKTimesNum 在一个数组中，其他数都出现了M次，只有一个数字出现了K次
func GetOnlyKTimesNum(arr []int, k int, m int) int {
	var t = [32]int{}
	for i := 0; i < len(arr); i++ {
		for j := 0; j < 32; j++ {
			if ((arr[i] >> j) & 1) != 0 {
				t[j]++
			}
		}
	}

	var answer = 0
	for i := 0; i < 32; i++ {
		if t[i]%m != 0 {
			answer |= 1 << i
		}
	}
	return answer
}
