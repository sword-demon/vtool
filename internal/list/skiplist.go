package list

import (
	"cmp"
	"errors"
	"math/rand"
	"time"
)

// SkipNode 跳表节点
type SkipNode[T cmp.Ordered] struct {
	Value T
	Next  []*SkipNode[T] // 每一层的后继节点指针
}

// SkipList 跳表
type SkipList[T cmp.Ordered] struct {
	head     *SkipNode[T]
	rand     *rand.Rand
	maxLevel int
	length   int
}

// NewSkipList 创建新的跳表
func NewSkipList[T cmp.Ordered]() *SkipList[T] {
	// 创建头节点，使用最大层数（实际使用时会动态调整）
	maxLevel := 16 // 最大16层
	head := &SkipNode[T]{
		Value: *new(T), // 哨兵节点，零值
		Next:  make([]*SkipNode[T], maxLevel),
	}

	return &SkipList[T]{
		head:     head,
		maxLevel: 0, // 从0层开始
		length:   0,
		rand:     rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// Insert 插入元素
func (s *SkipList[T]) Insert(value T) {
	// 查找插入位置和更新路径
	update := make([]*SkipNode[T], s.maxLevel+1)
	current := s.head

	// 从最高层开始查找插入位置
	for level := s.maxLevel; level >= 0; level-- {
		for current.Next[level] != nil && current.Next[level].Value < value {
			current = current.Next[level]
		}
		update[level] = current
	}

	// 如果值已存在，不插入
	if current.Next[0] != nil && current.Next[0].Value == value {
		return
	}

	// 随机决定层数
	newLevel := s.randomLevel()

	// 确保update数组有足够的空间
	if newLevel > s.maxLevel {
		// 重新分配update数组
		newUpdate := make([]*SkipNode[T], newLevel+1)
		copy(newUpdate, update)
		update = newUpdate

		for level := s.maxLevel + 1; level <= newLevel; level++ {
			update[level] = s.head
		}
		s.maxLevel = newLevel
	}

	// 创建新节点
	newNode := &SkipNode[T]{
		Value: value,
		Next:  make([]*SkipNode[T], newLevel+1),
	}

	// 在每一层插入节点
	for level := 0; level <= newLevel; level++ {
		newNode.Next[level] = update[level].Next[level]
		update[level].Next[level] = newNode
	}

	s.length++
}

// Search 查找元素
func (s *SkipList[T]) Search(value T) bool {
	current := s.head

	// 从最高层开始查找
	for level := s.maxLevel; level >= 0; level-- {
		for current.Next[level] != nil && current.Next[level].Value < value {
			current = current.Next[level]
		}
	}

	// 检查下一个节点是否为目标值
	current = current.Next[0]
	return current != nil && current.Value == value
}

// Remove 删除元素
func (s *SkipList[T]) Remove(value T) bool {
	update := make([]*SkipNode[T], s.maxLevel+1)
	current := s.head

	// 查找要删除的节点
	for level := s.maxLevel; level >= 0; level-- {
		for current.Next[level] != nil && current.Next[level].Value < value {
			current = current.Next[level]
		}
		update[level] = current
	}

	current = current.Next[0]
	if current == nil || current.Value != value {
		return false // 未找到
	}

	// 从每一层中删除节点
	for level := 0; level <= s.maxLevel; level++ {
		if update[level].Next[level] != current {
			break // 已经不在当前层
		}
		update[level].Next[level] = current.Next[level]
	}

	// 减少层数（如果最高层为空）
	for s.maxLevel > 0 && s.head.Next[s.maxLevel] == nil {
		s.maxLevel--
	}

	s.length--
	return true
}

// Contains 检查元素是否存在
func (s *SkipList[T]) Contains(value T) bool {
	return s.Search(value)
}

// Size 返回跳表中的元素数量
func (s *SkipList[T]) Size() int {
	return s.length
}

// IsEmpty 检查跳表是否为空
func (s *SkipList[T]) IsEmpty() bool {
	return s.length == 0
}

// ToSlice 转换为切片（已排序）
func (s *SkipList[T]) ToSlice() []T {
	result := make([]T, s.length)
	current := s.head.Next[0]
	for i := 0; i < s.length; i++ {
		result[i] = current.Value
		current = current.Next[0]
	}
	return result
}

// Clear 清空跳表
func (s *SkipList[T]) Clear() {
	maxLevel := 16
	s.head = &SkipNode[T]{
		Value: *new(T), // 零值
		Next:  make([]*SkipNode[T], maxLevel),
	}
	s.maxLevel = 0
	s.length = 0
}

// Min 返回最小值
func (s *SkipList[T]) Min() (T, error) {
	if s.length == 0 {
		var zero T
		return zero, errors.New("list is empty")
	}
	return s.head.Next[0].Value, nil
}

// Max 返回最大值
func (s *SkipList[T]) Max() (T, error) {
	if s.length == 0 {
		var zero T
		return zero, errors.New("list is empty")
	}
	current := s.head.Next[0]
	for current.Next[0] != nil {
		current = current.Next[0]
	}
	return current.Value, nil
}

// randomLevel 随机生成层数
func (s *SkipList[T]) randomLevel() int {
	level := 0
	// 50%的概率进入下一层
	for level < 16 && s.rand.Float64() < 0.5 {
		level++
	}
	return level
}
