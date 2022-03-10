package class03

type Stack struct {
	size int
	arr  [100]int
}

func PopStack(stack *Stack) int {
	if nil == stack || stack.size == 0 {
		return -1
	}
	num := stack.arr[stack.size-1]
	stack.size--
	return num
}
func PutStack(stack *Stack, val int) {
	if nil == stack {
		return
	}

	stack.arr[stack.size] = val
	stack.size++
}

func StackSize(stack Stack) int {
	return stack.size
}
