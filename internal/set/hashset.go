package set

// HashSet 基于map实现的哈希集合
type HashSet[T comparable] struct {
	items map[T]struct{} // 使用空结构体节省内存
}

// NewHashSet 创建新的HashSet
func NewHashSet[T comparable]() *HashSet[T] {
	return &HashSet[T]{
		items: make(map[T]struct{}),
	}
}

// Add 添加元素到集合
func (s *HashSet[T]) Add(item T) {
	s.items[item] = struct{}{}
}

// Remove 从集合中删除元素
func (s *HashSet[T]) Remove(item T) {
	delete(s.items, item)
}

// Contains 检查元素是否在集合中
func (s *HashSet[T]) Contains(item T) bool {
	_, exists := s.items[item]
	return exists
}

// Size 返回集合大小
func (s *HashSet[T]) Size() int {
	return len(s.items)
}

// IsEmpty 检查集合是否为空
func (s *HashSet[T]) IsEmpty() bool {
	return len(s.items) == 0
}

// Clear 清空集合
func (s *HashSet[T]) Clear() {
	for k := range s.items {
		delete(s.items, k)
	}
}

// ToSlice 返回集合的所有元素为切片
func (s *HashSet[T]) ToSlice() []T {
	items := make([]T, 0, len(s.items))
	for item := range s.items {
		items = append(items, item)
	}
	return items
}

// Union 返回两个集合的并集
func (s *HashSet[T]) Union(other *HashSet[T]) *HashSet[T] {
	result := NewHashSet[T]()

	// 添加当前集合的元素
	for item := range s.items {
		result.Add(item)
	}

	// 添加另一个集合的元素
	for item := range other.items {
		result.Add(item)
	}

	return result
}

// Intersect 返回两个集合的交集
func (s *HashSet[T]) Intersect(other *HashSet[T]) *HashSet[T] {
	result := NewHashSet[T]()

	// 遍历较小的集合
	if s.Size() < other.Size() {
		for item := range s.items {
			if other.Contains(item) {
				result.Add(item)
			}
		}
	} else {
		for item := range other.items {
			if s.Contains(item) {
				result.Add(item)
			}
		}
	}

	return result
}

// Difference 返回两个集合的差集 (s - other)
func (s *HashSet[T]) Difference(other *HashSet[T]) *HashSet[T] {
	result := NewHashSet[T]()

	for item := range s.items {
		if !other.Contains(item) {
			result.Add(item)
		}
	}

	return result
}