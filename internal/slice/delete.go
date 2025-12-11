package slice

import "errors"

// Delete 删除index位置的元素
// index 范围在 [0, len(src)-1]
func Delete[T any](src []T, index int) ([]T, error) {
	length := len(src)
	if index < 0 || index >= length {
		return nil, errors.New("index is out of range")
	}

	// 从index开始，所有元素向前移动一位
	for i := index; i < length-1; i++ {
		src[i] = src[i+1]
	}
	// 去掉最后一个元素
	return src[:length-1], nil
}