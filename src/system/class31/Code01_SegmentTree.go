package class31

type SegmentTree struct {
	maxLen int
	arr    []int
	sum    []int
	lazy   []int
	change []int
	modify []bool
}

func (segmentTree *SegmentTree) Init(origin []int) {
	segmentTree.maxLen = len(origin) + 1
	segmentTree.arr = make([]int, segmentTree.maxLen+1)
	for i := 1; i < segmentTree.maxLen; i++ {
		segmentTree.arr[i] = origin[i-1]
	}
	segmentTree.sum = make([]int, segmentTree.maxLen<<2)
	segmentTree.lazy = make([]int, segmentTree.maxLen<<2)
	segmentTree.change = make([]int, segmentTree.maxLen<<2)
	segmentTree.modify = make([]bool, segmentTree.maxLen<<2)

	// 构建初始树
	segmentTree.build(1, segmentTree.maxLen-1, 1)
}

func (segmentTree SegmentTree) build(left int, right int, root int) {
	if left == right {
		segmentTree.sum[root] = segmentTree.arr[left]
		return
	}
	mid := left + ((right - left) >> 1)
	segmentTree.build(left, mid, root<<1)
	segmentTree.build(mid+1, right, root<<1|1)

	segmentTree.pushUp(root)
}

func (segmentTree SegmentTree) pushUp(root int) {
	segmentTree.sum[root] = segmentTree.sum[root<<1] + segmentTree.sum[root<<1|1]
}

func (segmentTree SegmentTree) pushDown(root int, leftNum int, rightNum int) {

	// check update
	if segmentTree.modify[root] {
		segmentTree.modify[root<<1] = true
		segmentTree.modify[root<<1|1] = true

		segmentTree.change[root<<1] = segmentTree.change[root]
		segmentTree.change[root<<1|1] = segmentTree.change[root]

		segmentTree.sum[root<<1] = segmentTree.change[root] * leftNum
		segmentTree.sum[root<<1|1] = segmentTree.change[root] * rightNum

		segmentTree.lazy[root<<1] = 0
		segmentTree.lazy[root<<1|1] = 0

		segmentTree.modify[root] = false
	}
	// check lazy add
	if segmentTree.lazy[root] != 0 {
		segmentTree.sum[root<<1] += segmentTree.lazy[root] * leftNum
		segmentTree.lazy[root<<1] += segmentTree.lazy[root]

		segmentTree.sum[root<<1|1] += segmentTree.lazy[root] * rightNum
		segmentTree.lazy[root<<1|1] += segmentTree.lazy[root]

		segmentTree.lazy[root] = 0
	}
}

func (segmentTree SegmentTree) Add(start int, end int, val int, left int, right int, root int) {

	// base case
	if start <= left && right <= end {
		segmentTree.sum[root] += (right - left + 1) * val
		segmentTree.lazy[root] += val
		return
	}

	mid := left + ((right - left) >> 1)

	segmentTree.pushDown(root, mid-left+1, right-mid)
	if start <= mid {
		segmentTree.Add(start, end, val, left, mid, root<<1)
	}
	if end > mid {
		segmentTree.Add(start, end, val, mid+1, right, root<<1|1)
	}
	segmentTree.pushUp(root)
}

func (segmentTree SegmentTree) Update(start int, end int, val int, left int, right int, root int) {
	if start <= left && right <= end {
		segmentTree.sum[root] = (right - left + 1) * val
		segmentTree.change[root] = val
		segmentTree.lazy[root] = 0
		segmentTree.modify[root] = true
		return
	}
	mid := left + ((right - left) >> 1)
	segmentTree.pushDown(root, mid-left+1, right-mid)

	if start <= mid {
		segmentTree.Update(start, end, val, left, mid, root<<1)
	}
	if end > mid {
		segmentTree.Update(start, end, val, mid+1, right, root<<1|1)
	}
	segmentTree.pushUp(root)
}

func (segmentTree SegmentTree) Query(start int, end int, left int, right int, root int) int {
	if start <= left && right <= end {
		return segmentTree.sum[root]
	}
	mid := left + ((right - left) >> 1)
	segmentTree.pushDown(root, mid-left+1, right-mid)
	ans := 0
	if start <= mid {
		ans += segmentTree.Query(start, end, left, mid, root<<1)
	}
	if end > mid {
		ans += segmentTree.Query(start, end, mid+1, right, root<<1|1)
	}

	return ans
}
