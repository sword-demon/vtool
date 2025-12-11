package mapx

// LinkedNode 链表节点
type LinkedNode[K comparable, V any] struct {
	Key   K
	Value V
	Prev  *LinkedNode[K, V]
	Next  *LinkedNode[K, V]
}

// LinkedMap 保持插入顺序的映射
type LinkedMap[K comparable, V any] struct {
	items  map[K]*LinkedNode[K, V]
	head   *LinkedNode[K, V]
	tail   *LinkedNode[K, V]
	length int
}

// NewLinkedMap 创建新的LinkedMap
func NewLinkedMap[K comparable, V any]() *LinkedMap[K, V] {
	return &LinkedMap[K, V]{
		items:  make(map[K]*LinkedNode[K, V]),
		head:   nil,
		tail:   nil,
		length: 0,
	}
}

// Put 添加或更新键值对
func (lm *LinkedMap[K, V]) Put(key K, value V) {
	if node, exists := lm.items[key]; exists {
		// 键已存在，更新值并移到尾部
		node.Value = value
		if lm.tail != node {
			lm.moveToTail(node)
		}
	} else {
		// 新键，创建新节点并加到尾部
		newNode := &LinkedNode[K, V]{
			Key:   key,
			Value: value,
		}
		lm.items[key] = newNode
		lm.addToTail(newNode)
		lm.length++
	}
}

// Get 获取值
func (lm *LinkedMap[K, V]) Get(key K) (V, bool) {
	if node, exists := lm.items[key]; exists {
		return node.Value, true
	}
	var zero V
	return zero, false
}

// Remove 删除键值对
func (lm *LinkedMap[K, V]) Remove(key K) {
	if node, exists := lm.items[key]; exists {
		lm.removeNode(node)
		delete(lm.items, key)
		lm.length--
	}
}

// Contains 检查键是否存在
func (lm *LinkedMap[K, V]) Contains(key K) bool {
	_, exists := lm.items[key]
	return exists
}

// Size 返回键值对数量
func (lm *LinkedMap[K, V]) Size() int {
	return lm.length
}

// IsEmpty 检查是否为空
func (lm *LinkedMap[K, V]) IsEmpty() bool {
	return lm.length == 0
}

// Clear 清空映射
func (lm *LinkedMap[K, V]) Clear() {
	lm.items = make(map[K]*LinkedNode[K, V])
	lm.head = nil
	lm.tail = nil
	lm.length = 0
}

// Keys 返回所有键（按插入顺序）
func (lm *LinkedMap[K, V]) Keys() []K {
	keys := make([]K, 0, lm.length)
	current := lm.head
	for current != nil {
		keys = append(keys, current.Key)
		current = current.Next
	}
	return keys
}

// Values 返回所有值（按插入顺序）
func (lm *LinkedMap[K, V]) Values() []V {
	values := make([]V, 0, lm.length)
	current := lm.head
	for current != nil {
		values = append(values, current.Value)
		current = current.Next
	}
	return values
}

// Entries 返回所有键值对（按插入顺序）
func (lm *LinkedMap[K, V]) Entries() []struct {
	K K
	V V
} {
	entries := make([]struct {
		K K
		V V
	}, 0, lm.length)
	current := lm.head
	for current != nil {
		entries = append(entries, struct {
			K K
			V V
		}{K: current.Key, V: current.Value})
		current = current.Next
	}
	return entries
}

// First 返回第一个键值对
func (lm *LinkedMap[K, V]) First() (K, V, bool) {
	if lm.head == nil {
		var zeroK K
		var zeroV V
		return zeroK, zeroV, false
	}
	return lm.head.Key, lm.head.Value, true
}

// Last 返回最后一个键值对
func (lm *LinkedMap[K, V]) Last() (K, V, bool) {
	if lm.tail == nil {
		var zeroK K
		var zeroV V
		return zeroK, zeroV, false
	}
	return lm.tail.Key, lm.tail.Value, true
}

// ToMap 转换为Go内置map
func (lm *LinkedMap[K, V]) ToMap() map[K]V {
	result := make(map[K]V, lm.length)
	for k, node := range lm.items {
		result[k] = node.Value
	}
	return result
}

// addToTail 添加节点到尾部
func (lm *LinkedMap[K, V]) addToTail(node *LinkedNode[K, V]) {
	if lm.tail == nil {
		// 链表为空
		lm.head = node
		lm.tail = node
	} else {
		// 添加到尾部
		node.Prev = lm.tail
		node.Next = nil
		lm.tail.Next = node
		lm.tail = node
	}
}

// removeNode 移除节点
func (lm *LinkedMap[K, V]) removeNode(node *LinkedNode[K, V]) {
	if node.Prev != nil {
		node.Prev.Next = node.Next
	} else {
		// 是头节点
		lm.head = node.Next
	}

	if node.Next != nil {
		node.Next.Prev = node.Prev
	} else {
		// 是尾节点
		lm.tail = node.Prev
	}
}

// moveToTail 将节点移到尾部
func (lm *LinkedMap[K, V]) moveToTail(node *LinkedNode[K, V]) {
	if lm.tail == node {
		return // 已经是尾节点
	}

	// 从当前位置移除
	if node.Prev != nil {
		node.Prev.Next = node.Next
	} else {
		// 是头节点
		lm.head = node.Next
	}
	if node.Next != nil {
		node.Next.Prev = node.Prev
	}

	// 添加到尾部
	node.Prev = lm.tail
	node.Next = nil
	if lm.tail != nil {
		lm.tail.Next = node
	}
	lm.tail = node
}
