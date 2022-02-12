package class03

type Queue struct {
	begin int
	end   int
	size  int
	arr   [100]int
}

func PutQueue(queue *Queue, val int) {
	if nil == queue {
		return
	}
	queue.begin--
	if queue.begin == -1 {
		queue.begin = 99
	}

	queue.arr[queue.end] = val
	queue.size++
}

func PopQueue(queue *Queue) int {

	num := queue.arr[queue.end]
	queue.size++

	return num
}

func QueueSize(queue *Queue) int {
	return queue.size
}
