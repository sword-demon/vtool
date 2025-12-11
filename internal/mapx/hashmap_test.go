package mapx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashMap(t *testing.T) {
	t.Run("基本操作", func(t *testing.T) {
		m := NewHashMap[string, int]()

		// 测试初始状态
		assert.True(t, m.IsEmpty())
		assert.Equal(t, 0, m.Size())

		// 添加键值对
		m.Put("apple", 1)
		m.Put("banana", 2)
		m.Put("cherry", 3)

		assert.False(t, m.IsEmpty())
		assert.Equal(t, 3, m.Size())

		// 获取值
		val, exists := m.Get("apple")
		assert.True(t, exists)
		assert.Equal(t, 1, val)

		val, exists = m.Get("banana")
		assert.True(t, exists)
		assert.Equal(t, 2, val)

		// 更新值
		m.Put("apple", 10)
		val, exists = m.Get("apple")
		assert.True(t, exists)
		assert.Equal(t, 10, val)

		// 检查键存在
		assert.True(t, m.Contains("cherry"))
		assert.False(t, m.Contains("grape"))

		// 删除键
		m.Remove("banana")
		assert.Equal(t, 2, m.Size())
		assert.False(t, m.Contains("banana"))
		assert.True(t, m.Contains("apple"))
		assert.True(t, m.Contains("cherry"))

		// 清空
		m.Clear()
		assert.True(t, m.IsEmpty())
		assert.Equal(t, 0, m.Size())
	})

	t.Run("迭代顺序", func(t *testing.T) {
		m := NewHashMap[string, int]()

		// 按顺序插入
		m.Put("a", 1)
		m.Put("b", 2)
		m.Put("c", 3)

		// 验证键的顺序
		keys := m.Keys()
		assert.Equal(t, []string{"a", "b", "c"}, keys)

		// 验证值的顺序
		values := m.Values()
		assert.Equal(t, []int{1, 2, 3}, values)

		// 验证Entries
		entries := m.Entries()
		assert.Equal(t, 3, len(entries))
		assert.Equal(t, "a", entries[0].K)
		assert.Equal(t, 1, entries[0].V)
		assert.Equal(t, "b", entries[1].K)
		assert.Equal(t, 2, entries[1].V)
		assert.Equal(t, "c", entries[2].K)
		assert.Equal(t, 3, entries[2].V)
	})

	t.Run("更新后顺序保持", func(t *testing.T) {
		m := NewHashMap[string, int]()

		m.Put("a", 1)
		m.Put("b", 2)
		m.Put("c", 3)

		// 更新已有键的值
		m.Put("b", 20)

		// 验证键顺序不变
		keys := m.Keys()
		assert.Equal(t, []string{"a", "b", "c"}, keys)

		// 验证值正确
		val, _ := m.Get("b")
		assert.Equal(t, 20, val)
	})

	t.Run("整数键", func(t *testing.T) {
		m := NewHashMap[int, string]()

		m.Put(1, "one")
		m.Put(2, "two")
		m.Put(3, "three")

		assert.Equal(t, 3, m.Size())

		val, exists := m.Get(2)
		assert.True(t, exists)
		assert.Equal(t, "two", val)

		keys := m.Keys()
		assert.Equal(t, []int{1, 2, 3}, keys)
	})
}