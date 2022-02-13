package class21

import "math"

func BobLiveProbability1(N, M, row, col, K int) float64 {

	live := processBobLiveProbability(N, M, row, col, K)

	return float64(live) / math.Pow(float64(4), float64(K))
}

func processBobLiveProbability(N, M, row, col, rest int) int {
	if row < 0 || col < 0 || row >= N || col >= M {
		return 0
	}
	if rest == 0 {
		return 1
	}
	live := 0
	live += processBobLiveProbability(N, M, row-1, col, rest-1)
	live += processBobLiveProbability(N, M, row+1, col, rest-1)
	live += processBobLiveProbability(N, M, row, col-1, rest-1)
	live += processBobLiveProbability(N, M, row, col+1, rest-1)

	return live
}

func BobLiveProbability2(N, M, row, col, K int) float64 {

	dp := make([][][]int, N)
	for i := 0; i < N; i++ {
		dp[i] = make([][]int, M)
		for j := 0; j < M; j++ {
			dp[i][j] = make([]int, K+1)
			dp[i][j][0] = 1
		}
	}
	for rest := 1; rest <= K; rest++ {
		for row := 0; row < N; row++ {
			for col := 0; col < M; col++ {
				live := pick(dp, N, M, row-1, col, rest-1)
				live += pick(dp, N, M, row+1, col, rest-1)
				live += pick(dp, N, M, row, col-1, rest-1)
				live += pick(dp, N, M, row, col+1, rest-1)

				dp[row][col][rest] = live
			}
		}
	}

	return float64(dp[row][col][K]) / math.Pow(float64(4), float64(K))
}
func pick(dp [][][]int, N, M, row, col, rest int) int {
	if row < 0 || col < 0 || row >= N || col >= M {
		return 0
	}

	return dp[row][col][rest]
}
