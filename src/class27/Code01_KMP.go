package class27

func GetIndexOf(haystack string, needle string) int {

	if len(haystack) < len(needle) {
		return -1
	}

	str1 := []byte(haystack)
	str2 := []byte(needle)

	next := getNextArray(str2)

	x := 0
	y := 0

	for x < len(str1) && y < len(str2) {
		if str1[x] == str2[y] {
			x++
			y++
		} else if next[y] != -1 {
			y = next[y]
		} else {
			x++
		}
	}

	if y == len(str2) {
		return x - y
	}

	return -1
}

func getNextArray(str2 []byte) []int {
	if len(str2) == 0 {
		return make([]int, 0)
	}

	if len(str2) == 1 {
		return []int{-1}
	}

	if len(str2) == 2 {
		return []int{-1, 0}
	}
	next := make([]int, len(str2))
	next[0] = -1
	next[1] = 0

	// 当前正在计算 i 位置的值
	i := 2

	// 当前是哪个位置的值在和 `i-1` 位置的值做比较
	cn := 0

	for i < len(str2) {
		if str2[i-1] == str2[cn] {
			cn++
			next[i] = cn
			i++
		} else if cn > 0 {
			cn = next[cn]
		} else {
			next[i] = 0
			i++
		}
	}

	return next
}
