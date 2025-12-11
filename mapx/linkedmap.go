package mapx

import "github.com/sword-demon/vtool/internal/mapx"

// LinkedNode 链表节点
type LinkedNode[K comparable, V any] = mapx.LinkedNode[K, V]

// LinkedMap 保持插入顺序的映射
type LinkedMap[K comparable, V any] = mapx.LinkedMap[K, V]

// NewLinkedMap 创建新的LinkedMap
func NewLinkedMap[K comparable, V any]() *LinkedMap[K, V] {
	return mapx.NewLinkedMap[K, V]()
}
