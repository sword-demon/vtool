package list

import (
	"cmp"

	"github.com/sword-demon/vtool/internal/list"
)

// ArrayList 动态数组
type ArrayList[T cmp.Ordered] = list.ArrayList[T]

// NewArrayList 创建新的ArrayList
func NewArrayList[T cmp.Ordered]() *ArrayList[T] {
	return list.NewArrayList[T]()
}

// NewArrayListWithCapacity 创建指定初始容量的ArrayList
func NewArrayListWithCapacity[T cmp.Ordered](capacity int) *ArrayList[T] {
	return list.NewArrayListWithCapacity[T](capacity)
}
