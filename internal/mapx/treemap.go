package mapx

import (
	"cmp"
	"errors"
)

// TreeNode 树节点
type TreeNode[K cmp.Ordered, V any] struct {
	Key    K
	Value  V
	Left   *TreeNode[K, V]
	Right  *TreeNode[K, V]
	Parent *TreeNode[K, V]
	Color  bool // true = Red, false = Black
}

// TreeMap 有序映射，基于二叉搜索树
type TreeMap[K cmp.Ordered, V any] struct {
	root   *TreeNode[K, V]
	length int
}

// NewTreeMap 创建新的TreeMap
func NewTreeMap[K cmp.Ordered, V any]() *TreeMap[K, V] {
	return &TreeMap[K, V]{
		root:   nil,
		length: 0,
	}
}

// Put 添加或更新键值对
func (tm *TreeMap[K, V]) Put(key K, value V) {
	tm.root = tm.insert(tm.root, key, value, nil)
	tm.length++
}

// Get 获取值
func (tm *TreeMap[K, V]) Get(key K) (V, bool) {
	node := tm.search(tm.root, key)
	if node == nil {
		var zero V
		return zero, false
	}
	return node.Value, true
}

// Remove 删除键
func (tm *TreeMap[K, V]) Remove(key K) {
	deleted := tm.delete(tm.root, key)
	if deleted {
		tm.length--
	}
}

// Contains 检查键是否存在
func (tm *TreeMap[K, V]) Contains(key K) bool {
	return tm.search(tm.root, key) != nil
}

// Size 返回键值对数量
func (tm *TreeMap[K, V]) Size() int {
	return tm.length
}

// IsEmpty 检查是否为空
func (tm *TreeMap[K, V]) IsEmpty() bool {
	return tm.length == 0
}

// Clear 清空映射
func (tm *TreeMap[K, V]) Clear() {
	tm.root = nil
	tm.length = 0
}

// Keys 返回所有键（按排序顺序）
func (tm *TreeMap[K, V]) Keys() []K {
	var keys []K
	tm.inOrderTraversal(tm.root, &keys)
	return keys
}

// Values 返回所有值（按排序顺序）
func (tm *TreeMap[K, V]) Values() []V {
	var values []V
	tm.inOrderTraversalValues(tm.root, &values)
	return values
}

// Entries 返回所有键值对（按排序顺序）
func (tm *TreeMap[K, V]) Entries() []struct {
	K K
	V V
} {
	var entries []struct {
		K K
		V V
	}
	tm.inOrderTraversalEntries(tm.root, &entries)
	return entries
}

// Min 返回最小键
func (tm *TreeMap[K, V]) Min() (K, error) {
	if tm.root == nil {
		var zero K
		return zero, errors.New("tree map is empty")
	}
	node := tm.min(tm.root)
	return node.Key, nil
}

// Max 返回最大键
func (tm *TreeMap[K, V]) Max() (K, error) {
	if tm.root == nil {
		var zero K
		return zero, errors.New("tree map is empty")
	}
	node := tm.max(tm.root)
	return node.Key, nil
}

// insert 插入节点
func (tm *TreeMap[K, V]) insert(node *TreeNode[K, V], key K, value V, parent *TreeNode[K, V]) *TreeNode[K, V] {
	if node == nil {
		return &TreeNode[K, V]{
			Key:    key,
			Value:  value,
			Parent: parent,
			Color:  true, // 新节点默认为红色
		}
	}

	if key < node.Key {
		node.Left = tm.insert(node.Left, key, value, node)
	} else if key > node.Key {
		node.Right = tm.insert(node.Right, key, value, node)
	} else {
		// 键已存在，更新值
		node.Value = value
		tm.length-- // 抵消后面的length++
	}

	return node
}

// search 搜索节点
func (tm *TreeMap[K, V]) search(node *TreeNode[K, V], key K) *TreeNode[K, V] {
	if node == nil {
		return nil
	}

	if key < node.Key {
		return tm.search(node.Left, key)
	} else if key > node.Key {
		return tm.search(node.Right, key)
	} else {
		return node
	}
}

// delete 删除节点
func (tm *TreeMap[K, V]) delete(node *TreeNode[K, V], key K) bool {
	if node == nil {
		return false
	}

	if key < node.Key {
		return tm.delete(node.Left, key)
	} else if key > node.Key {
		return tm.delete(node.Right, key)
	} else {
		// 找到要删除的节点
		if node.Left == nil && node.Right == nil {
			// 叶子节点
			if node.Parent != nil {
				if node.Parent.Left == node {
					node.Parent.Left = nil
				} else {
					node.Parent.Right = nil
				}
			}
			return true
		} else if node.Left == nil {
			// 只有右子节点
			if node.Parent != nil {
				if node.Parent.Left == node {
					node.Parent.Left = node.Right
				} else {
					node.Parent.Right = node.Right
				}
			}
			node.Right.Parent = node.Parent
			return true
		} else if node.Right == nil {
			// 只有左子节点
			if node.Parent != nil {
				if node.Parent.Left == node {
					node.Parent.Left = node.Left
				} else {
					node.Parent.Right = node.Left
				}
			}
			node.Left.Parent = node.Parent
			return true
		} else {
			// 有两个子节点，找到后继节点
			successor := tm.min(node.Right)
			node.Key = successor.Key
			node.Value = successor.Value
			return tm.delete(node.Right, successor.Key)
		}
	}
}

// min 找到最小节点
func (tm *TreeMap[K, V]) min(node *TreeNode[K, V]) *TreeNode[K, V] {
	current := node
	for current.Left != nil {
		current = current.Left
	}
	return current
}

// max 找到最大节点
func (tm *TreeMap[K, V]) max(node *TreeNode[K, V]) *TreeNode[K, V] {
	current := node
	for current.Right != nil {
		current = current.Right
	}
	return current
}

// inOrderTraversal 中序遍历收集键
func (tm *TreeMap[K, V]) inOrderTraversal(node *TreeNode[K, V], keys *[]K) {
	if node != nil {
		tm.inOrderTraversal(node.Left, keys)
		*keys = append(*keys, node.Key)
		tm.inOrderTraversal(node.Right, keys)
	}
}

// inOrderTraversalValues 中序遍历收集值
func (tm *TreeMap[K, V]) inOrderTraversalValues(node *TreeNode[K, V], values *[]V) {
	if node != nil {
		tm.inOrderTraversalValues(node.Left, values)
		*values = append(*values, node.Value)
		tm.inOrderTraversalValues(node.Right, values)
	}
}

// inOrderTraversalEntries 中序遍历收集键值对
func (tm *TreeMap[K, V]) inOrderTraversalEntries(node *TreeNode[K, V], entries *[]struct {
	K K
	V V
}) {
	if node != nil {
		tm.inOrderTraversalEntries(node.Left, entries)
		*entries = append(*entries, struct {
			K K
			V V
		}{K: node.Key, V: node.Value})
		tm.inOrderTraversalEntries(node.Right, entries)
	}
}
