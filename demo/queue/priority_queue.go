package queue

// F8DoCompare 比较方法，用于比较 a 和 b 的大小
// a<b 返回 true，a>=b 返回 false。
type F8DoCompare[T any] func(a T, b T) bool

// s6PriorityQueue 优先队列，小顶堆实现
// 并发安全问题由外面的延迟队列管
type s6PriorityQueue[T any] struct {
	s5data      []T
	nowSize     int
	maxSize     int
	f8DoCompare F8DoCompare[T]
}

func F8NewS6PriorityQueue[T any](maxSize int, f8DoCompare F8DoCompare[T]) *s6PriorityQueue[T] {
	// 这里让下标从 1 开始，为了让小顶堆中的元素编号和数组下标对齐
	return &s6PriorityQueue[T]{
		s5data:      make([]T, 1, maxSize),
		nowSize:     1,
		maxSize:     maxSize,
		f8DoCompare: f8DoCompare,
	}
}

func (p7this *s6PriorityQueue[T]) f8IsFull() bool {
	return p7this.nowSize == p7this.maxSize
}

func (p7this *s6PriorityQueue[T]) f8IsEmpty() bool {
	return 1 >= p7this.nowSize
}

func (p7this *s6PriorityQueue[T]) f8Enqueue(data T) error {
	if p7this.f8IsFull() {
		return ErrQueueIsFull
	}
	// 把数据放到堆的末尾
	p7this.s5data = append(p7this.s5data, data)
	p7this.nowSize++
	// 调整堆使之重新成为小根堆
	childIndex := p7this.nowSize - 1
	parentIndex := (p7this.nowSize - 1) / 2
	for 0 < parentIndex && p7this.f8DoCompare(p7this.s5data[childIndex], p7this.s5data[parentIndex]) {
		p7this.s5data[childIndex], p7this.s5data[parentIndex] = p7this.s5data[parentIndex], p7this.s5data[childIndex]
		childIndex = parentIndex
		parentIndex = parentIndex / 2
	}

	return nil
}

func (p7this *s6PriorityQueue[T]) f8Dequeue() (T, error) {
	if p7this.f8IsEmpty() {
		var data T
		return data, ErrQueueIsEmpty
	}
	// 堆顶元素出堆
	data := p7this.s5data[1]
	// 把最后一个元素移动到堆顶
	p7this.s5data[1] = p7this.s5data[p7this.nowSize-1]
	p7this.s5data = p7this.s5data[:p7this.nowSize-1]
	p7this.nowSize--
	// 调整堆使之重新成为小根堆
	parentIndex := 1
	minIndex := parentIndex
	for {
		leftIndex := 2 * parentIndex
		if leftIndex <= p7this.nowSize-1 && p7this.f8DoCompare(p7this.s5data[leftIndex], p7this.s5data[minIndex]) {
			minIndex = leftIndex
		}
		rightIndex := 2*parentIndex + 1
		if rightIndex <= p7this.nowSize-1 && p7this.f8DoCompare(p7this.s5data[rightIndex], p7this.s5data[minIndex]) {
			minIndex = rightIndex
		}
		if minIndex == parentIndex {
			break
		}
		p7this.s5data[minIndex], p7this.s5data[parentIndex] = p7this.s5data[parentIndex], p7this.s5data[minIndex]
		parentIndex = minIndex
	}
	return data, nil
}

func (p7this *s6PriorityQueue[T]) f8GetTop() (T, error) {
	if p7this.f8IsEmpty() {
		var data T
		return data, ErrQueueIsEmpty
	}
	return p7this.s5data[1], nil
}
