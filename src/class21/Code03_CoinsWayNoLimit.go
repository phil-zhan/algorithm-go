package class21

func CoinsWayNoLimit(arr []int, aim int) int {

	if nil == arr || len(arr) == 0 || aim < 0 {
		return 0
	}

	return processCoinsWayNoLimit(arr, 0, aim)
}

func processCoinsWayNoLimit(arr []int, index int, rest int) int {

	if index == len(arr) {
		if rest == 0 {
			return 1
		} else {
			return 0
		}
	}
	ways := 0
	for zhang := 0; zhang*arr[index] <= rest; zhang++ {
		ways += processCoinsWayNoLimit(arr, index+1, rest-zhang*arr[index])
	}
	return ways
}

func CoinsWayNoLimitDp1(arr []int, aim int) int {
	if nil == arr || len(arr) == 0 || aim < 0 {
		return 0
	}
	N := len(arr)

	dp := make([][]int, N+1)
	for i := 0; i <= N; i++ {
		dp[i] = make([]int, aim+1)
	}
	dp[N][0] = 1
	for index := N - 1; index >= 0; index-- {
		for rest := 0; rest <= aim; rest++ {
			ways := 0
			for zhang := 0; zhang*arr[index] <= rest; zhang++ {
				ways += dp[index+1][rest-zhang*arr[index]]
			}
			dp[index][rest] = ways
		}
	}

	return dp[0][aim]
}

func CoinsWayNoLimitDp2(arr []int, aim int) int {
	if nil == arr || len(arr) == 0 || aim < 0 {
		return 0
	}
	N := len(arr)

	dp := make([][]int, N+1)
	for i := 0; i <= N; i++ {
		dp[i] = make([]int, aim+1)
	}
	dp[N][0] = 1
	for index := N - 1; index >= 0; index-- {
		for rest := 0; rest <= aim; rest++ {

			dp[index][rest] = dp[index+1][rest]
			if rest-arr[index] >= 0 {
				dp[index][rest] += dp[index][rest-arr[index]]
			}

		}
	}

	return dp[0][aim]
}
