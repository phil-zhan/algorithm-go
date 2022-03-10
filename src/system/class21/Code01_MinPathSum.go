package class21

import "math"

// IntMax 系统最大值
const IntMax = int(^uint(0) >> 1)

func MinPathSum1(grid [][]int) int {

	if grid == nil || len(grid) == 0 || grid[0] == nil || len(grid[0]) == 0 {
		return 0
	}

	return processMinPathSum1(grid, 0, 0)
}

func processMinPathSum1(grid [][]int, row, col int) int {

	N := len(grid)
	M := len(grid[0])
	if row == N-1 && col == M-1 {
		return grid[row][col]
	}

	left := IntMax
	if col < M-1 {
		left = grid[row][col] + processMinPathSum1(grid, row, col+1)
	}
	down := IntMax
	if row < N-1 {
		down = grid[row][col] + processMinPathSum1(grid, row+1, col)
	}

	return int(math.Min(float64(left), float64(down)))
}

func MinPathSumDp(grid [][]int) int {
	N := len(grid)
	M := len(grid[0])

	dp := make([][]int, N)
	for i := 0; i < N; i++ {
		dp[i] = make([]int, M)
	}

	dp[N-1][M-1] = grid[N-1][M-1]
	for row := N - 2; row >= 0; row-- {
		dp[row][M-1] = grid[row][M-1] + dp[row+1][M-1]
	}

	for col := M - 2; col >= 0; col-- {
		dp[N-1][col] = grid[N-1][col] + dp[N-1][col+1]
	}

	for row := N - 2; row >= 0; row-- {
		for col := M - 2; col >= 0; col-- {
			dp[row][col] = grid[row][col] + int(math.Min(float64(dp[row][col+1]), float64(dp[row+1][col])))
		}
	}

	return dp[0][0]

}

func MinPathSumDp2(grid [][]int) int {

	N := len(grid)
	M := len(grid[0])

	dp := make([]int, M)

	dp[M-1] = grid[N-1][M-1]
	for col := M - 2; col >= 0; col-- {
		dp[col] = grid[N-1][col] + dp[col+1]
	}

	for row := N - 2; row >= 0; row-- {
		dp[M-1] = dp[M-1] + grid[row][M-1]
		for col := M - 2; col >= 0; col-- {
			dp[col] = grid[row][col] + int(math.Min(float64(dp[col]), float64(dp[col+1])))
		}
	}

	return dp[0]

}
