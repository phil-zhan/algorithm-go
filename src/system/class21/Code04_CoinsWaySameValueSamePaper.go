package class21

func CoinsWaySameValueSamePaper(arr []int, aim int) int {

	map1 := make(map[int]int)

	for i := 0; i < len(arr); i++ {
		temp := map1[arr[i]]
		temp++
		map1[arr[i]] = temp
	}

	coins := make([]int, len(map1))
	zhangs := make([]int, len(map1))

	index := 0
	for key, value := range map1 {
		coins[index] = key
		zhangs[index] = value
		index++
	}

	return processCoinsWaySameValueSamePaper(coins, zhangs, 0, aim)
}

func processCoinsWaySameValueSamePaper(coins, zhangs []int, index int, rest int) int {

	if index == len(coins) {
		if rest == 0 {
			return 1
		} else {
			return 0
		}
	}
	ways := 0
	for zhang := 0; zhang <= zhangs[index] && zhang*coins[index] <= rest; zhang++ {
		ways += processCoinsWaySameValueSamePaper(coins, zhangs, index+1, rest-zhang*coins[index])
	}

	return ways
}

func CoinsWaySameValueSamePaperDp1(arr []int, aim int) int {

	map1 := make(map[int]int)

	for i := 0; i < len(arr); i++ {
		temp := map1[arr[i]]
		temp++
		map1[arr[i]] = temp
	}

	coins := make([]int, len(map1))
	zhangs := make([]int, len(map1))

	index := 0
	for key, value := range map1 {
		coins[index] = key
		zhangs[index] = value
		index++
	}

	N := len(coins)
	dp := make([][]int, N+1)
	for i := 0; i <= N; i++ {
		dp[i] = make([]int, aim+1)
	}

	dp[N][0] = 1

	for index := N - 1; index >= 0; index-- {
		for rest := 0; rest <= aim; rest++ {
			ways := 0
			for zhang := 0; zhang <= zhangs[index] && zhang*coins[index] <= rest; zhang++ {
				ways += dp[index+1][rest-zhang*coins[index]]
			}
			dp[index][rest] = ways
		}
	}
	return dp[0][aim]
}

func CoinsWaySameValueSamePaperDp2(arr []int, aim int) int {

	map1 := make(map[int]int)

	for i := 0; i < len(arr); i++ {
		temp := map1[arr[i]]
		temp++
		map1[arr[i]] = temp
	}

	coins := make([]int, len(map1))
	zhangs := make([]int, len(map1))

	index := 0
	for key, value := range map1 {
		coins[index] = key
		zhangs[index] = value
		index++
	}

	N := len(coins)
	dp := make([][]int, N+1)
	for i := 0; i <= N; i++ {
		dp[i] = make([]int, aim+1)
	}

	dp[N][0] = 1

	for index := N - 1; index >= 0; index-- {
		for rest := 0; rest <= aim; rest++ {
			dp[index][rest] = dp[index+1][rest]
			if rest-coins[index] >= 0 {
				dp[index][rest] += dp[index][rest-coins[index]]
			}
			if rest-coins[index]*(zhangs[index]+1) >= 0 {
				dp[index][rest] -= dp[index+1][rest-coins[index]*(zhangs[index]+1)]
			}
		}
	}
	return dp[0][aim]
}
