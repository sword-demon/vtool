package sets

import (
	"cmp"

	"github.com/sword-demon/vtool/internal/set"
)

// SortedSet 排序并去重切片中的元素
func SortedSet[T cmp.Ordered](src []T) []T {
	return set.SortedSet(src)
}

// SortedSetDesc 降序排序并去重
func SortedSetDesc[T cmp.Ordered](src []T) []T {
	return set.SortedSetDesc(src)
}

// Unique 仅去重，保持原有顺序
func Unique[T comparable](src []T) []T {
	return set.Unique(src)
}
