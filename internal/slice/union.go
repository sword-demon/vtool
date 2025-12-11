package slice

// Union 求两个切片的并集，返回去重的结果
// 使用map来去重，保持元素的原始顺序
func Union[T comparable](src1, src2 []T) []T {
	// 如果两个切片都为空，返回空切片
	if len(src1) == 0 && len(src2) == 0 {
		return []T{}
	}

	seen := make(map[T]bool)
	var result []T

	// 添加第一个切片的元素
	for _, item := range src1 {
		if !seen[item] {
			seen[item] = true
			result = append(result, item)
		}
	}

	// 添加第二个切片的元素，去重
	for _, item := range src2 {
		if !seen[item] {
			seen[item] = true
			result = append(result, item)
		}
	}

	return result
}
