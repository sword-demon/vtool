package list

import (
	"cmp"
	"errors"
)

// Node 双向链表节点
type Node[T cmp.Ordered] struct {
	Value T
	Prev  *Node[T]
	Next  *Node[T]
}

// LinkedList 双向链表
type LinkedList[T cmp.Ordered] struct {
	head   *Node[T]
	tail   *Node[T]
	length int
}

// NewLinkedList 创建新的双向链表
func NewLinkedList[T cmp.Ordered]() *LinkedList[T] {
	return &LinkedList[T]{}
}

// Add 添加元素到链表末尾
func (l *LinkedList[T]) Add(value T) {
	l.Insert(l.length, value)
}

// Insert 在指定位置插入元素
func (l *LinkedList[T]) Insert(index int, value T) error {
	if index < 0 || index > l.length {
		return errors.New("index out of range")
	}

	newNode := &Node[T]{Value: value}

	if l.length == 0 {
		// 空链表
		l.head = newNode
		l.tail = newNode
	} else if index == 0 {
		// 插入到头部
		newNode.Next = l.head
		l.head.Prev = newNode
		l.head = newNode
	} else if index == l.length {
		// 插入到尾部
		newNode.Prev = l.tail
		l.tail.Next = newNode
		l.tail = newNode
	} else {
		// 插入到中间
		current := l.getNode(index)
		newNode.Prev = current.Prev
		newNode.Next = current
		current.Prev.Next = newNode
		current.Prev = newNode
	}

	l.length++
	return nil
}

// Remove 删除指定位置的元素
func (l *LinkedList[T]) Remove(index int) error {
	if l.length == 0 {
		return errors.New("list is empty")
	}
	if index < 0 || index >= l.length {
		return errors.New("index out of range")
	}

	if l.length == 1 {
		// 只有一个元素
		l.head = nil
		l.tail = nil
	} else if index == 0 {
		// 删除头部
		l.head = l.head.Next
		l.head.Prev = nil
	} else if index == l.length-1 {
		// 删除尾部
		l.tail = l.tail.Prev
		l.tail.Next = nil
	} else {
		// 删除中间元素
		current := l.getNode(index)
		current.Prev.Next = current.Next
		current.Next.Prev = current.Prev
	}

	l.length--
	return nil
}

// Get 获取指定位置的元素
func (l *LinkedList[T]) Get(index int) (T, error) {
	if l.length == 0 {
		var zero T
		return zero, errors.New("list is empty")
	}
	if index < 0 || index >= l.length {
		var zero T
		return zero, errors.New("index out of range")
	}

	node := l.getNode(index)
	return node.Value, nil
}

// Set 设置指定位置的元素值
func (l *LinkedList[T]) Set(index int, value T) error {
	if l.length == 0 {
		return errors.New("list is empty")
	}
	if index < 0 || index >= l.length {
		return errors.New("index out of range")
	}

	node := l.getNode(index)
	node.Value = value
	return nil
}

// IndexOf 查找元素第一次出现的位置
func (l *LinkedList[T]) IndexOf(value T) int {
	current := l.head
	for i := 0; i < l.length; i++ {
		if current.Value == value {
			return i
		}
		current = current.Next
	}
	return -1
}

// Contains 检查元素是否存在
func (l *LinkedList[T]) Contains(value T) bool {
	return l.IndexOf(value) != -1
}

// Size 返回链表长度
func (l *LinkedList[T]) Size() int {
	return l.length
}

// IsEmpty 检查链表是否为空
func (l *LinkedList[T]) IsEmpty() bool {
	return l.length == 0
}

// Clear 清空链表
func (l *LinkedList[T]) Clear() {
	l.head = nil
	l.tail = nil
	l.length = 0
}

// ToSlice 转换为切片
func (l *LinkedList[T]) ToSlice() []T {
	result := make([]T, l.length)
	current := l.head
	for i := 0; i < l.length; i++ {
		result[i] = current.Value
		current = current.Next
	}
	return result
}

// FromSlice 从切片创建链表
func (l *LinkedList[T]) FromSlice(values []T) {
	l.Clear()
	for _, value := range values {
		l.Add(value)
	}
}

// getNode 获取指定位置的节点（内部方法）
func (l *LinkedList[T]) getNode(index int) *Node[T] {
	var current *Node[T]
	if index < l.length/2 {
		// 从头部开始遍历
		current = l.head
		for i := 0; i < index; i++ {
			current = current.Next
		}
	} else {
		// 从尾部开始遍历
		current = l.tail
		for i := l.length - 1; i > index; i-- {
			current = current.Prev
		}
	}
	return current
}