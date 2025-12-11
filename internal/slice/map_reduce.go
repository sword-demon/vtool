package slice

// Map 将每个元素映射为新值，返回新的切片
// mapper 是一个函数，接受原元素并返回新元素
func Map[T any, R any](src []T, mapper func(T) R) []R {
	result := make([]R, len(src))
	for i, item := range src {
		result[i] = mapper(item)
	}
	return result
}

// Reduce 将切片的所有元素聚合为一个值
// reducer 是一个函数，接受累计值和当前元素，返回新的累计值
// initial 是初始累计值
func Reduce[T any, R any](src []T, reducer func(R, T) R, initial R) R {
	result := initial
	for _, item := range src {
		result = reducer(result, item)
	}
	return result
}

// Filter 过滤元素，返回满足条件的新切片
// predicate 是一个函数，接受元素并返回bool
func Filter[T any](src []T, predicate func(T) bool) []T {
	var result []T
	for _, item := range src {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}