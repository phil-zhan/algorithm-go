package class02

// Swap 不使用额外变量，交换两个数
func Swap(a *int, b *int) {

	// 两个数一样，就不要再交换了，不然会导致结果为0
	if *a == *b {
		return
	}

	*a = (*a) ^ (*b)
	*b = (*a) ^ (*b)
	*a = (*a) ^ (*b)
}
