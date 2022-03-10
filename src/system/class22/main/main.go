package main

import (
	class222 "algorithm-go/src/system/class22"
	"fmt"
)

func main() {
	number := class222.SplitNumber(2)
	number2 := class222.SplitNumberDP1(2)
	number3 := class222.SplitNumberDP2(2)
	fmt.Println(number)
	fmt.Println(number2)
	fmt.Println(number3)

	limit1 := class222.MinCoinsNoLimit1([]int{23, 16, 21, 17, 28, 4, 27}, 40)
	limit2 := class222.MinCoinsNoLimit2([]int{23, 16, 21, 17, 28, 4, 27}, 40)
	limit3 := class222.MinCoinsNoLimit3([]int{23, 16, 21, 17, 28, 4, 27}, 40)
	fmt.Println(limit1)
	fmt.Println(limit2)
	fmt.Println(limit3)

	monster1 := class222.KillMonster(5, 4, 9)
	monster2 := class222.KillMonsterDp1(5, 4, 9)
	monster3 := class222.KillMonsterDp2(5, 4, 9)
	fmt.Println(monster1)
	fmt.Println(monster2)
	fmt.Println(monster3)
}
