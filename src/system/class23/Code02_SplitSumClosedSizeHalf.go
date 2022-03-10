package class23

import "math"

func SplitSumClosedSizeHalf(arr []int) int {
	if nil == arr || len(arr) < 2 {
		return 0
	}
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	if (len(arr) & 1) == 0 {
		return processSizeHalf(arr, 0, len(arr)>>1, sum>>1)
	} else {
		return int(math.Max(float64(processSizeHalf(arr, 0, len(arr)>>1, sum>>1)), float64(processSizeHalf(arr, 0, (len(arr)>>1)+1, sum>>1))))
	}
}

func processSizeHalf(arr []int, cur int, pick int, rest int) int {
	if cur == len(arr) {
		if pick == 0 {
			return 0
		}
		return -1
	}
	p1 := processSizeHalf(arr, cur+1, pick, rest)

	p2 := -1
	next := -1
	if arr[cur] <= rest {
		next = processSizeHalf(arr, cur+1, pick-1, rest-arr[cur])
	}
	if next != -1 {
		p2 = arr[cur] + next
	}
	return int(math.Max(float64(p1), float64(p2)))
}

func SplitSumClosedSizeHalfDP(arr []int) int {
	if arr == nil || len(arr) < 2 {
		return 0
	}
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	sum /= 2
	N := len(arr)
	M := (N + 1) / 2
	dp := make([][][]int, N+1)
	//int[][][] dp = new int[N + 1][M + 1][sum + 1];
	for i := 0; i <= N; i++ {
		dp[i] = make([][]int, M+1)
		for j := 0; j <= M; j++ {
			dp[i][j] = make([]int, sum+1)
			for k := 0; k <= sum; k++ {
				dp[i][j][k] = -1
			}
		}
	}
	for rest := 0; rest <= sum; rest++ {
		dp[N][0][rest] = 0
	}
	for i := N - 1; i >= 0; i-- {
		for picks := 0; picks <= M; picks++ {
			for rest := 0; rest <= sum; rest++ {
				p1 := dp[i+1][picks][rest]
				// 就是要使用arr[i]这个数
				p2 := -1
				next := -1
				if picks-1 >= 0 && arr[i] <= rest {
					next = dp[i+1][picks-1][rest-arr[i]]
				}
				if next != -1 {
					p2 = arr[i] + next
				}
				dp[i][picks][rest] = int(math.Max(float64(p1), float64(p2)))
			}
		}
	}
	if (len(arr) & 1) == 0 {
		return dp[0][len(arr)/2][sum]
	} else {
		return int(math.Max(float64(dp[0][len(arr)/2][sum]), float64(dp[0][(len(arr)/2)+1][sum])))
	}
}
