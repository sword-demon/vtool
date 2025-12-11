package mapx

import "github.com/sword-demon/vtool/internal/mapx"

// HashMap 基于Go内置map实现的增强版映射
type HashMap[K comparable, V any] = mapx.HashMap[K, V]

// NewHashMap 创建新的HashMap
func NewHashMap[K comparable, V any]() *HashMap[K, V] {
	return mapx.NewHashMap[K, V]()
}
