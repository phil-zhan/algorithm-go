package class23

import "math"

// NQueens n皇后问题
// n*n的矩阵上摆放皇后，不能同行、不能同列、也不能同斜线
func NQueens(n int) int {
	if n < 1 {
		return 0
	}
	record := make([]int, n)

	return processNQueens(0, record, n)
}
func processNQueens(row int, record []int, n int) int {
	if row == n {
		return 1
	}

	res := 0
	for col := 0; col < n; col++ {
		if isValid(record, row, col) {
			record[row] = col
			res += processNQueens(row+1, record, n)
		}
	}

	return res
}

func isValid(record []int, row int, col int) bool {
	for i := 0; i < row; i++ {
		// 原始记录 i行  record[i]列
		if col == record[i] || math.Abs(float64(row-i)) == math.Abs(float64(col-record[i])) {
			return false
		}
	}
	return true
}

// NQueens2
//位运算版本
func NQueens2(n int) int {
	if n < 1 {
		return 0
	}
	limit := 0
	if n == 32 {
		limit = -1
	} else {
		limit = (1 << n) - 1
	}
	return processNQueens2(limit, 0, 0, 0)
}
func processNQueens2(limit int, colLimit, leftDownLimit, rightDownLimit int) int {
	if limit == colLimit {
		return 1
	}

	pos := limit & (^(colLimit | leftDownLimit | rightDownLimit))
	res := 0
	for pos != 0 {
		mostRightOne := pos & (^pos + 1)
		pos -= mostRightOne
		res += processNQueens2(limit, colLimit|mostRightOne, (leftDownLimit|mostRightOne)<<1, (rightDownLimit|mostRightOne)>>1)
	}

	return res
}
