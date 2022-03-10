package class22

import "math"

// IntMax 系统最大值
const IntMax = int(^uint(0) >> 1)

// MinCoinsNoLimit1 递归版本
// 从左往右的尝试模型
func MinCoinsNoLimit1(arr []int, aim int) int {
	ans := processMinCoins(arr, 0, aim)
	if ans == IntMax {
		return -1
	}
	return ans
}

func processMinCoins(arr []int, index int, rest int) int {
	if index == len(arr) {
		if rest == 0 {
			return 0
		}
		return IntMax
	}

	ans := IntMax
	for zhang := 0; arr[index]*zhang <= rest; zhang++ {
		next := processMinCoins(arr, index+1, rest-arr[index]*zhang)
		if next != IntMax {
			ans = int(math.Min(float64(ans), float64(next+zhang)))
		}
	}

	return ans
}

// MinCoinsNoLimit2 常规动态规划
func MinCoinsNoLimit2(arr []int, aim int) int {
	N := len(arr)
	dp := make([][]int, N+1)
	for i := 0; i <= N; i++ {
		dp[i] = make([]int, aim+1)
	}

	dp[N][0] = 0
	for i := 1; i <= aim; i++ {
		dp[N][i] = IntMax
	}

	for index := N - 1; index >= 0; index-- {
		for rest := 0; rest <= aim; rest++ {

			ans := IntMax
			for zhang := 0; arr[index]*zhang <= rest; zhang++ {
				next := dp[index+1][rest-arr[index]*zhang]
				if next != IntMax {
					ans = int(math.Min(float64(ans), float64(next+zhang)))
				}
			}

			dp[index][rest] = ans
		}
	}

	if dp[0][aim] == IntMax {
		return -1
	}

	return dp[0][aim]
}

// MinCoinsNoLimit3 动态规划。斜率优化版本
func MinCoinsNoLimit3(arr []int, aim int) int {
	N := len(arr)
	dp := make([][]int, N+1)
	for i := 0; i <= N; i++ {
		dp[i] = make([]int, aim+1)
	}

	dp[N][0] = 0
	for i := 1; i <= aim; i++ {
		dp[N][i] = IntMax
	}

	for index := N - 1; index >= 0; index-- {
		for rest := 0; rest <= aim; rest++ {

			dp[index][rest] = dp[index+1][rest]
			// 不加此条件  可能会导致最大值越界
			// && dp[index][rest-arr[index]] != IntMax
			if rest-arr[index] >= 0 && dp[index][rest-arr[index]] != IntMax {
				dp[index][rest] = int(math.Min(float64(dp[index][rest]), float64(dp[index][rest-arr[index]]+1)))
			}
		}
	}

	if dp[0][aim] == IntMax {
		return -1
	}

	return dp[0][aim]
}
