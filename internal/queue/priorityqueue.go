package queue

import "errors"

// PriorityItem 优先级队列中的元素
type PriorityItem[T any] struct {
	Value    T
	Priority int // 优先级，数值越小优先级越高
}

// PriorityQueue 优先级队列 - 基于最小堆实现
type PriorityQueue[T any] struct {
	items []PriorityItem[T]
}

// NewPriorityQueue 创建新的优先级队列
func NewPriorityQueue[T any]() *PriorityQueue[T] {
	return &PriorityQueue[T]{
		items: make([]PriorityItem[T], 0),
	}
}

// Enqueue 入队 - 添加元素和优先级
func (pq *PriorityQueue[T]) Enqueue(value T, priority int) {
	item := PriorityItem[T]{
		Value:    value,
		Priority: priority,
	}
	pq.items = append(pq.items, item)
	pq.siftUp(len(pq.items) - 1)
}

// Dequeue 出队 - 移除并返回优先级最高的元素
func (pq *PriorityQueue[T]) Dequeue() (T, int, error) {
	if pq.IsEmpty() {
		var zero T
		return zero, 0, errors.New("priority queue is empty")
	}

	// 取出根节点（优先级最高的）
	item := pq.items[0]
	value := item.Value
	priority := item.Priority

	// 将最后一个元素移到根节点
	lastIndex := len(pq.items) - 1
	pq.items[0] = pq.items[lastIndex]
	pq.items = pq.items[:lastIndex]

	// 如果还有元素，重新堆化
	if len(pq.items) > 0 {
		pq.siftDown(0)
	}

	return value, priority, nil
}

// Peek 查看优先级最高的元素
func (pq *PriorityQueue[T]) Peek() (T, int, error) {
	if pq.IsEmpty() {
		var zero T
		return zero, 0, errors.New("priority queue is empty")
	}

	return pq.items[0].Value, pq.items[0].Priority, nil
}

// Size 返回队列中的元素数量
func (pq *PriorityQueue[T]) Size() int {
	return len(pq.items)
}

// IsEmpty 检查队列是否为空
func (pq *PriorityQueue[T]) IsEmpty() bool {
	return len(pq.items) == 0
}

// Clear 清空队列
func (pq *PriorityQueue[T]) Clear() {
	pq.items = pq.items[:0]
}

// ToSlice 转换为切片（按优先级排序）
func (pq *PriorityQueue[T]) ToSlice() []PriorityItem[T] {
	result := make([]PriorityItem[T], len(pq.items))
	copy(result, pq.items)
	return result
}

// siftUp 向上堆化（插入时使用）
func (pq *PriorityQueue[T]) siftUp(index int) {
	for index > 0 {
		parent := (index - 1) / 2

		// 如果当前元素优先级更高（数值更小），交换
		if pq.items[index].Priority < pq.items[parent].Priority {
			pq.items[index], pq.items[parent] = pq.items[parent], pq.items[index]
			index = parent
		} else {
			break
		}
	}
}

// siftDown 向下堆化（删除时使用）
func (pq *PriorityQueue[T]) siftDown(index int) {
	length := len(pq.items)
	for {
		left := 2*index + 1
		right := 2*index + 2
		smallest := index

		// 找到子节点中优先级更高的
		if left < length && pq.items[left].Priority < pq.items[smallest].Priority {
			smallest = left
		}
		if right < length && pq.items[right].Priority < pq.items[smallest].Priority {
			smallest = right
		}

		// 如果根节点已经是优先级最高的，停止
		if smallest == index {
			break
		}

		// 交换元素
		pq.items[index], pq.items[smallest] = pq.items[smallest], pq.items[index]
		index = smallest
	}
}