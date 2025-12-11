package slice

import "github.com/sword-demon/vtool/internal/slice"

// Find 查找第一个满足条件的元素
// predicate 是一个函数，接受元素并返回bool
// 返回找到的元素索引和值，如果没找到返回-1
func Find[Src any](src []Src, predicate func(Src) bool) (int, Src, bool) {
	index, val, found := slice.Find(src, predicate)
	return index, val, found
}