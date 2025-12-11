package queue

import "github.com/sword-demon/vtool/internal/queue"

// Queue 普通队列 - FIFO (先进先出)
type Queue[T any] = queue.Queue[T]

// NewQueue 创建新的队列
func NewQueue[T any]() *Queue[T] {
	return queue.NewQueue[T]()
}

// PriorityItem 优先级队列中的元素
type PriorityItem[T any] = queue.PriorityItem[T]

// PriorityQueue 优先级队列
type PriorityQueue[T any] = queue.PriorityQueue[T]

// NewPriorityQueue 创建新的优先级队列
func NewPriorityQueue[T any]() *PriorityQueue[T] {
	return queue.NewPriorityQueue[T]()
}
