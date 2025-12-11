package list

import (
	"cmp"
	"errors"
)

// ArrayList 动态数组实现，类似于Java的ArrayList
type ArrayList[T cmp.Ordered] struct {
	items    []T
	length   int
	capacity int
}

// NewArrayList 创建新的ArrayList
func NewArrayList[T cmp.Ordered]() *ArrayList[T] {
	return &ArrayList[T]{
		items:    make([]T, 0, 10), // 默认初始容量10
		length:   0,
		capacity: 10,
	}
}

// NewArrayListWithCapacity 创建指定初始容量的ArrayList
func NewArrayListWithCapacity[T cmp.Ordered](capacity int) *ArrayList[T] {
	if capacity < 1 {
		capacity = 1
	}
	return &ArrayList[T]{
		items:    make([]T, 0, capacity),
		length:   0,
		capacity: capacity,
	}
}

// Add 添加元素到列表末尾
func (l *ArrayList[T]) Add(value T) {
	l.Insert(l.length, value)
}

// Insert 在指定位置插入元素
func (l *ArrayList[T]) Insert(index int, value T) error {
	if index < 0 || index > l.length {
		return errors.New("index out of range")
	}

	// 检查是否需要扩容
	if l.length >= l.capacity {
		l.expand()
	}

	// 使用append来插入元素，保持items的len和cap正确
	if index == l.length {
		// 插入到末尾
		l.items = append(l.items, value)
	} else {
		// 插入到中间
		l.items = append(l.items[:index+1], l.items[index:]...)
		l.items[index] = value
	}

	l.length++
	return nil
}

// Remove 删除指定位置的元素
func (l *ArrayList[T]) Remove(index int) error {
	if l.length == 0 {
		return errors.New("list is empty")
	}
	if index < 0 || index >= l.length {
		return errors.New("index out of range")
	}

	// 使用 append 来删除元素
	l.items = append(l.items[:index], l.items[index+1:]...)
	l.length--

	// 检查是否需要收缩容量
	l.maybeShrink()

	return nil
}

// Get 获取指定位置的元素
func (l *ArrayList[T]) Get(index int) (T, error) {
	if l.length == 0 {
		var zero T
		return zero, errors.New("list is empty")
	}
	if index < 0 || index >= l.length {
		var zero T
		return zero, errors.New("index out of range")
	}

	return l.items[index], nil
}

// Set 设置指定位置的元素值
func (l *ArrayList[T]) Set(index int, value T) error {
	if l.length == 0 {
		return errors.New("list is empty")
	}
	if index < 0 || index >= l.length {
		return errors.New("index out of range")
	}

	l.items[index] = value
	return nil
}

// IndexOf 查找元素第一次出现的位置
func (l *ArrayList[T]) IndexOf(value T) int {
	for i := 0; i < l.length; i++ {
		if l.items[i] == value {
			return i
		}
	}
	return -1
}

// Contains 检查元素是否存在
func (l *ArrayList[T]) Contains(value T) bool {
	return l.IndexOf(value) != -1
}

// Size 返回列表长度
func (l *ArrayList[T]) Size() int {
	return l.length
}

// IsEmpty 检查列表是否为空
func (l *ArrayList[T]) IsEmpty() bool {
	return l.length == 0
}

// Clear 清空列表
func (l *ArrayList[T]) Clear() {
	l.items = l.items[:0] // 重置len，保留cap
	l.length = 0
}

// ToSlice 转换为切片
func (l *ArrayList[T]) ToSlice() []T {
	result := make([]T, l.length)
	copy(result, l.items[:l.length])
	return result
}

// Capacity 返回当前容量
func (l *ArrayList[T]) Capacity() int {
	return l.capacity
}

// Trim 收缩容量以匹配当前大小
func (l *ArrayList[T]) Trim() {
	l.capacity = l.length
	l.items = l.items[:l.length]
}

// expand 扩容（内部方法）
func (l *ArrayList[T]) expand() {
	// 扩容为当前容量的2倍
	newCapacity := l.capacity * 2
	newItems := make([]T, 0, newCapacity)
	// 复制现有元素
	newItems = append(newItems, l.items...)
	l.items = newItems
	l.capacity = newCapacity
}

// maybeShrink 检查是否需要收缩容量（内部方法）
func (l *ArrayList[T]) maybeShrink() {
	// 当元素数量小于容量的1/4时，收缩容量
	if l.length > 0 && l.length < l.capacity/4 {
		newCapacity := l.capacity / 2
		if newCapacity < 10 {
			newCapacity = 10 // 最小容量
		}
		newItems := make([]T, 0, newCapacity)
		newItems = append(newItems, l.items...)
		l.items = newItems
		l.capacity = newCapacity
	}
}

// EnsureCapacity 确保容量至少为指定大小
func (l *ArrayList[T]) EnsureCapacity(minCapacity int) {
	if minCapacity > l.capacity {
		l.capacity = minCapacity
		newItems := make([]T, 0, l.capacity)
		newItems = append(newItems, l.items...)
		l.items = newItems
	}
}

// RemoveValue 删除第一个匹配的元素（值匹配）
func (l *ArrayList[T]) RemoveValue(value T) bool {
	index := l.IndexOf(value)
	if index == -1 {
		return false
	}
	l.Remove(index)
	return true
}
