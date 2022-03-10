package class22

func SplitNumber(n int) int {
	if n < 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	return process(1, n)
}

func process(pre, rest int) int {
	if rest == 0 {
		return 1
	}
	if rest < pre {
		return 0
	}

	ways := 0
	for i := pre; i <= rest; i++ {
		ways += process(i, rest-i)
	}

	return ways
}

func SplitNumberDP1(n int) int {
	if n < 0 {
		return 0
	}

	if n == 1 {
		return 1
	}
	dp := make([][]int, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+1)
		dp[i][0] = 1
		dp[i][i] = 1
	}

	for i := n - 1; i >= 1; i-- {
		for j := i + 1; j <= n; j++ {

			ways := 0
			for k := i; k <= j; k++ {
				ways += dp[i][j-k]
			}

			dp[i][j] = ways
		}
	}

	return dp[1][n]
}

func SplitNumberDP2(n int) int {
	if n < 0 {
		return 0
	}

	if n == 1 {
		return 1
	}
	dp := make([][]int, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+1)
		dp[i][0] = 1
		dp[i][i] = 1
	}

	for i := n - 1; i >= 1; i-- {
		for j := i + 1; j <= n; j++ {

			dp[i][j] = dp[i+1][j]
			dp[i][j] += dp[i][j-i]
		}
	}

	return dp[1][n]
}
