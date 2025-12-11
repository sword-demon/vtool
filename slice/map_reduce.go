package slice

import "github.com/sword-demon/vtool/internal/slice"

// Map 将每个元素映射为新值，返回新的切片
func Map[Src any, Dst any](src []Src, mapper func(Src) Dst) []Dst {
	return slice.Map(src, mapper)
}

// Reduce 将切片的所有元素聚合为一个值
func Reduce[Src any, R any](src []Src, reducer func(R, Src) R, initial R) R {
	return slice.Reduce(src, reducer, initial)
}

// Filter 过滤元素，返回满足条件的新切片
func Filter[Src any](src []Src, predicate func(Src) bool) []Src {
	return slice.Filter(src, predicate)
}
