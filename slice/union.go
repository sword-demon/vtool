package slice

import "github.com/sword-demon/vtool/internal/slice"

// Union 求两个切片的并集，返回去重的结果
func Union[Src comparable](src1, src2 []Src) []Src {
	return slice.Union(src1, src2)
}
