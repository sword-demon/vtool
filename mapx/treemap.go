package mapx

import (
	"cmp"

	"github.com/sword-demon/vtool/internal/mapx"
)

// TreeNode 树节点
type TreeNode[K cmp.Ordered, V any] = mapx.TreeNode[K, V]

// TreeMap 有序映射，基于二叉搜索树
type TreeMap[K cmp.Ordered, V any] = mapx.TreeMap[K, V]

// NewTreeMap 创建新的TreeMap
func NewTreeMap[K cmp.Ordered, V any]() *TreeMap[K, V] {
	return mapx.NewTreeMap[K, V]()
}
