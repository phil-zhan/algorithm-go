package class23

import (
	"math"
)

// SplitSumClosed
// 给定一个正数数组arr，
//请把arr中所有的数分成两个集合，尽量让两个集合的累加和接近
//返回：
//最接近的情况下，较小集合的累加和
func SplitSumClosed(arr []int) int {
	if nil == arr || len(arr) < 2 {
		return 0
	}
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}

	return process(arr, 0, sum>>1)
}

func process(arr []int, cur int, rest int) int {
	if cur == len(arr) {
		return 0
	} else {
		p1 := process(arr, cur+1, rest)

		p2 := 0
		if arr[cur] <= rest {
			p2 = arr[cur] + process(arr, cur+1, rest-arr[cur])
		}
		return int(math.Max(float64(p1), float64(p2)))
	}
}

func SplitSumClosedDP(arr []int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	dp := make([][]int, len(arr)+1)
	for i := 0; i <= len(arr); i++ {
		dp[i] = make([]int, (sum>>1)+1)
	}

	for cur := len(arr) - 1; cur >= 0; cur-- {
		for rest := 0; rest <= (sum >> 1); rest++ {
			p1 := dp[cur+1][rest]

			p2 := 0
			if arr[cur] <= rest {
				p2 = arr[cur] + dp[cur+1][rest-arr[cur]]
			}
			dp[cur][rest] = int(math.Max(float64(p1), float64(p2)))
		}
	}

	return dp[0][(sum >> 1)]
}
