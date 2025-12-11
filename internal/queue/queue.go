package queue

import (
	"errors"
)

// Queue 普通队列 - FIFO (先进先出)
type Queue[T any] struct {
	items []T
}

// NewQueue 创建新的队列
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		items: make([]T, 0),
	}
}

// Enqueue 入队 - 添加元素到队列尾部
func (q *Queue[T]) Enqueue(value T) {
	q.items = append(q.items, value)
}

// Dequeue 出队 - 从队列头部移除并返回元素
// 如果队列为空，返回零值和错误
func (q *Queue[T]) Dequeue() (T, error) {
	if q.IsEmpty() {
		var zero T
		return zero, errors.New("queue is empty")
	}

	value := q.items[0]
	q.items = q.items[1:]
	return value, nil
}

// Peek 查看队首元素 - 返回队首元素但不移除
// 如果队列为空，返回零值和错误
func (q *Queue[T]) Peek() (T, error) {
	if q.IsEmpty() {
		var zero T
		return zero, errors.New("queue is empty")
	}

	return q.items[0], nil
}

// Size 返回队列中的元素数量
func (q *Queue[T]) Size() int {
	return len(q.items)
}

// IsEmpty 检查队列是否为空
func (q *Queue[T]) IsEmpty() bool {
	return len(q.items) == 0
}

// Clear 清空队列
func (q *Queue[T]) Clear() {
	q.items = q.items[:0]
}

// ToSlice 转换为切片
func (q *Queue[T]) ToSlice() []T {
	result := make([]T, len(q.items))
	copy(result, q.items)
	return result
}