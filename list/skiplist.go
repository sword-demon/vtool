package list

import (
	"cmp"

	"github.com/sword-demon/vtool/internal/list"
)

// SkipList 跳表
type SkipList[T cmp.Ordered] = list.SkipList[T]

// NewSkipList 创建新的跳表
func NewSkipList[T cmp.Ordered]() *SkipList[T] {
	return list.NewSkipList[T]()
}
