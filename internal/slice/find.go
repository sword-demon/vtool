package slice

// Find 查找第一个满足条件的元素
// predicate 是一个函数，接受元素并返回bool
// 返回找到的元素索引和值，如果没找到返回-1
func Find[T any](src []T, predicate func(T) bool) (int, T, bool) {
	for i, item := range src {
		if predicate(item) {
			return i, item, true
		}
	}
	var zeroValue T
	return -1, zeroValue, false
}