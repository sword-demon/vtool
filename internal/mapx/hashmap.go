package mapx

// HashMap 基于Go内置map实现的增强版映射
// 提供可预测的迭代顺序
type HashMap[K comparable, V any] struct {
	items map[K]V
	keys  []K // 维护插入顺序
}

// NewHashMap 创建新的HashMap
func NewHashMap[K comparable, V any]() *HashMap[K, V] {
	return &HashMap[K, V]{
		items: make(map[K]V),
		keys:  make([]K, 0),
	}
}

// Put 添加或更新键值对
func (m *HashMap[K, V]) Put(key K, value V) {
	if _, exists := m.items[key]; !exists {
		// 新键，添加到keys列表
		m.keys = append(m.keys, key)
	}
	m.items[key] = value
}

// Get 获取值
func (m *HashMap[K, V]) Get(key K) (V, bool) {
	val, exists := m.items[key]
	return val, exists
}

// Remove 删除键值对
func (m *HashMap[K, V]) Remove(key K) {
	if _, exists := m.items[key]; exists {
		delete(m.items, key)
		// 从keys列表中移除
		for i, k := range m.keys {
			if k == key {
				m.keys = append(m.keys[:i], m.keys[i+1:]...)
				break
			}
		}
	}
}

// Contains 检查键是否存在
func (m *HashMap[K, V]) Contains(key K) bool {
	_, exists := m.items[key]
	return exists
}

// Size 返回键值对数量
func (m *HashMap[K, V]) Size() int {
	return len(m.items)
}

// IsEmpty 检查是否为空
func (m *HashMap[K, V]) IsEmpty() bool {
	return len(m.items) == 0
}

// Clear 清空映射
func (m *HashMap[K, V]) Clear() {
	m.items = make(map[K]V)
	m.keys = m.keys[:0]
}

// Keys 返回所有键（按插入顺序）
func (m *HashMap[K, V]) Keys() []K {
	result := make([]K, len(m.keys))
	copy(result, m.keys)
	return result
}

// Values 返回所有值（按插入顺序）
func (m *HashMap[K, V]) Values() []V {
	result := make([]V, 0, len(m.items))
	for _, key := range m.keys {
		result = append(result, m.items[key])
	}
	return result
}

// ToMap 转换为Go内置map
func (m *HashMap[K, V]) ToMap() map[K]V {
	result := make(map[K]V, len(m.items))
	for k, v := range m.items {
		result[k] = v
	}
	return result
}

// Entries 返回所有键值对（按插入顺序）
func (m *HashMap[K, V]) Entries() []struct {
	K K
	V V
} {
	result := make([]struct {
		K K
		V V
	}, 0, len(m.items))
	for _, key := range m.keys {
		result = append(result, struct {
			K K
			V V
		}{K: key, V: m.items[key]})
	}
	return result
}