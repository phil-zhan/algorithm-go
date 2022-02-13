package class22

import "math"

func KillMonster(hp int, M int, times int) float64 {

	live := processKillMonster(times, M, hp)

	return 1 - (float64(live) / math.Pow(float64(M+1), float64(times)))
}

func processKillMonster(restTimes int, M int, restHp int) int {
	if restHp <= 0 {
		return 0
	}

	if restTimes == 0 {
		if restHp > 0 {
			return 1
		} else {
			return 0
		}
	}
	live := 0
	for i := 0; i <= M; i++ {
		live += processKillMonster(restTimes-1, M, restHp-i)
	}

	return live
}

func KillMonsterDp1(hp int, M int, times int) float64 {

	dp := make([][]int, times+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, hp+1)
	}

	for i := 1; i <= hp; i++ {
		dp[0][i] = 1
	}

	for restTimes := 1; restTimes <= times; restTimes++ {

		// 第一列永远是0
		for restHp := 1; restHp <= hp; restHp++ {

			live := 0
			for i := 0; i <= M; i++ {
				if restHp-i >= 0 {
					live += dp[restTimes-1][restHp-i]
				}
			}

			dp[restTimes][restHp] = live

		}
	}

	return 1 - (float64(dp[times][hp]) / math.Pow(float64(M+1), float64(times)))
}

func KillMonsterDp2(hp int, M int, times int) float64 {

	dp := make([][]int, times+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, hp+1)
	}

	for i := 1; i <= hp; i++ {
		dp[0][i] = 1
	}

	for i := 1; i <= hp; i++ {
		dp[0][i] = 1
	}

	for restTimes := 1; restTimes <= times; restTimes++ {
		// 第一列永远是0
		for restHp := 1; restHp <= hp; restHp++ {

			dp[restTimes][restHp] = dp[restTimes][restHp-1] + dp[restTimes-1][restHp]
			if (restHp - M - 1) >= 0 {
				dp[restTimes][restHp] -= dp[restTimes-1][restHp-M-1]
			}

		}
	}

	return 1 - (float64(dp[times][hp]) / math.Pow(float64(M+1), float64(times)))
}
