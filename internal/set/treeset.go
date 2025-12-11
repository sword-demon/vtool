package set

import (
	"cmp"
	"sort"
)

// TreeSet 基于排序切片实现的有序集合
// T必须是可比较且可排序的类型
type TreeSet[T cmp.Ordered] struct {
	items []T
}

// NewTreeSet 创建新的TreeSet
func NewTreeSet[T cmp.Ordered]() *TreeSet[T] {
	return &TreeSet[T]{
		items: make([]T, 0),
	}
}

// Add 添加元素到集合，自动保持有序
func (s *TreeSet[T]) Add(item T) {
	// 如果元素已存在，不添加
	if s.Contains(item) {
		return
	}

	// 插入元素并保持有序
	s.items = append(s.items, item)
	sort.Slice(s.items, func(i, j int) bool {
		return s.items[i] < s.items[j]
	})
}

// Remove 从集合中删除元素
func (s *TreeSet[T]) Remove(item T) {
	index := sort.Search(len(s.items), func(i int) bool {
		return s.items[i] >= item
	})

	// 找到精确匹配的元素
	if index < len(s.items) && s.items[index] == item {
		s.items = append(s.items[:index], s.items[index+1:]...)
	}
}

// Contains 检查元素是否在集合中
func (s *TreeSet[T]) Contains(item T) bool {
	index := sort.Search(len(s.items), func(i int) bool {
		return s.items[i] >= item
	})

	return index < len(s.items) && s.items[index] == item
}

// Size 返回集合大小
func (s *TreeSet[T]) Size() int {
	return len(s.items)
}

// IsEmpty 检查集合是否为空
func (s *TreeSet[T]) IsEmpty() bool {
	return len(s.items) == 0
}

// Clear 清空集合
func (s *TreeSet[T]) Clear() {
	s.items = s.items[:0]
}

// ToSlice 返回集合的所有元素为切片（已排序）
func (s *TreeSet[T]) ToSlice() []T {
	result := make([]T, len(s.items))
	copy(result, s.items)
	return result
}

// Union 返回两个集合的并集
func (s *TreeSet[T]) Union(other *TreeSet[T]) *TreeSet[T] {
	result := NewTreeSet[T]()

	// 合并两个有序切片
	i, j := 0, 0
	for i < len(s.items) && j < len(other.items) {
		if s.items[i] < other.items[j] {
			result.Add(s.items[i])
			i++
		} else if s.items[i] > other.items[j] {
			result.Add(other.items[j])
			j++
		} else {
			// 元素相等，只添加一次
			result.Add(s.items[i])
			i++
			j++
		}
	}

	// 添加剩余元素
	for i < len(s.items) {
		result.Add(s.items[i])
		i++
	}
	for j < len(other.items) {
		result.Add(other.items[j])
		j++
	}

	return result
}

// Intersect 返回两个集合的交集
func (s *TreeSet[T]) Intersect(other *TreeSet[T]) *TreeSet[T] {
	result := NewTreeSet[T]()

	// 双指针遍历两个有序切片
	i, j := 0, 0
	for i < len(s.items) && j < len(other.items) {
		if s.items[i] < other.items[j] {
			i++
		} else if s.items[i] > other.items[j] {
			j++
		} else {
			// 元素相等，加入结果
			result.Add(s.items[i])
			i++
			j++
		}
	}

	return result
}

// Difference 返回两个集合的差集 (s - other)
func (s *TreeSet[T]) Difference(other *TreeSet[T]) *TreeSet[T] {
	result := NewTreeSet[T]()

	// 双指针遍历
	i, j := 0, 0
	for i < len(s.items) && j < len(other.items) {
		if s.items[i] < other.items[j] {
			result.Add(s.items[i])
			i++
		} else if s.items[i] > other.items[j] {
			j++
		} else {
			// 元素相等，跳过
			i++
			j++
		}
	}

	// 添加s中剩余的元素
	for i < len(s.items) {
		result.Add(s.items[i])
		i++
	}

	return result
}
