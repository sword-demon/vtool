package sets

import (
	"cmp"

	"github.com/sword-demon/vtool/internal/set"
)

// TreeSet 基于排序切片实现的有序集合
type TreeSet[T cmp.Ordered] = set.TreeSet[T]

// NewTreeSet 创建新的TreeSet
func NewTreeSet[T cmp.Ordered]() *TreeSet[T] {
	return set.NewTreeSet[T]()
}
