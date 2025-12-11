package sets

import "github.com/sword-demon/vtool/internal/set"

// HashSet 基于map实现的哈希集合
type HashSet[T comparable] = set.HashSet[T]

// NewHashSet 创建新的HashSet
func NewHashSet[T comparable]() *HashSet[T] {
	return set.NewHashSet[T]()
}