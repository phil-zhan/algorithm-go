package class48

import (
	"math/rand"
	"time"
)

func FindKthLargest(nums []int, k int) int {
	return getKth(nums, 0, len(nums)-1, k-1)
}

func getKth(nums []int, left int, right int, index int) int {
	if left == right {
		return nums[index]
	}

	rand.Seed(time.Now().Unix())

	pivot := nums[left+rand.Intn(right-left+1)]

	ranges := partition(nums, left, right, pivot)

	if index >= ranges[0] && index <= ranges[1] {
		return nums[index]
	} else if index > ranges[1] {
		return getKth(nums, ranges[1]+1, right, index)
	} else {
		return getKth(nums, left, ranges[0]-1, index)
	}

}

func partition(nums []int, left int, right int, pivot int) []int {
	more := left - 1
	less := right + 1

	cur := left

	for cur < less {
		if nums[cur] > pivot {
			more++
			swap(&nums, more, cur)
			cur++
		} else if nums[cur] < pivot {
			less--
			swap(&nums, less, cur)
		} else {
			cur++
		}
	}

	return []int{more + 1, less - 1}
}

func swap(nums *[]int, i1 int, i2 int) {
	temp := (*nums)[i1]
	(*nums)[i1] = (*nums)[i2]
	(*nums)[i2] = temp
}
