package main

import (
	"algorithm-go/src/class22"
	"fmt"
)

func main() {
	number := class22.SplitNumber(2)
	number2 := class22.SplitNumberDP1(2)
	number3 := class22.SplitNumberDP2(2)
	fmt.Println(number)
	fmt.Println(number2)
	fmt.Println(number3)

	limit1 := class22.MinCoinsNoLimit1([]int{23, 16, 21, 17, 28, 4, 27}, 40)
	limit2 := class22.MinCoinsNoLimit2([]int{23, 16, 21, 17, 28, 4, 27}, 40)
	limit3 := class22.MinCoinsNoLimit3([]int{23, 16, 21, 17, 28, 4, 27}, 40)
	fmt.Println(limit1)
	fmt.Println(limit2)
	fmt.Println(limit3)

	monster1 := class22.KillMonster(5, 4, 9)
	monster2 := class22.KillMonsterDp1(5, 4, 9)
	monster3 := class22.KillMonsterDp2(5, 4, 9)
	fmt.Println(monster1)
	fmt.Println(monster2)
	fmt.Println(monster3)
}
