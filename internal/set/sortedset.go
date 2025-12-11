package set

import (
	"cmp"
	"sort"
)

// SortedSet 排序并去重切片中的元素
// 返回新的有序切片
// T必须是可比较且可排序的类型
func SortedSet[T cmp.Ordered](src []T) []T {
	if len(src) == 0 {
		return []T{}
	}

	// 复制切片避免修改原数据
	result := make([]T, len(src))
	copy(result, src)

	// 排序
	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})

	// 去重
	unique := make([]T, 0, len(result))
	for i, item := range result {
		if i == 0 || item != result[i-1] {
			unique = append(unique, item)
		}
	}

	return unique
}

// SortedSetDesc 降序排序并去重
func SortedSetDesc[T cmp.Ordered](src []T) []T {
	if len(src) == 0 {
		return []T{}
	}

	// 复制切片避免修改原数据
	result := make([]T, len(src))
	copy(result, src)

	// 降序排序
	sort.Slice(result, func(i, j int) bool {
		return result[i] > result[j]
	})

	// 去重
	unique := make([]T, 0, len(result))
	for i, item := range result {
		if i == 0 || item != result[i-1] {
			unique = append(unique, item)
		}
	}

	return unique
}

// Unique 仅去重，保持原有顺序
func Unique[T comparable](src []T) []T {
	if len(src) == 0 {
		return []T{}
	}

	seen := make(map[T]bool)
	result := make([]T, 0, len(src))

	for _, item := range src {
		if !seen[item] {
			seen[item] = true
			result = append(result, item)
		}
	}

	return result
}