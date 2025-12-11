package slice

import "github.com/sword-demon/vtool/internal/slice"

// Delete 在index处删除元素
// index 范围在 [0, len(src)-1]
func Delete[Src any](src []Src, index int) ([]Src, error) {
	res, err := slice.Delete[Src](src, index)
	return res, err
}