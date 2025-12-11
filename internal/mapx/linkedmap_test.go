package mapx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedMap(t *testing.T) {
	t.Run("基本操作", func(t *testing.T) {
		lm := NewLinkedMap[string, int]()

		// 测试初始状态
		assert.True(t, lm.IsEmpty())
		assert.Equal(t, 0, lm.Size())

		// 添加键值对
		lm.Put("apple", 1)
		lm.Put("banana", 2)
		lm.Put("cherry", 3)

		assert.False(t, lm.IsEmpty())
		assert.Equal(t, 3, lm.Size())

		// 获取值
		val, exists := lm.Get("apple")
		assert.True(t, exists)
		assert.Equal(t, 1, val)

		val, exists = lm.Get("banana")
		assert.True(t, exists)
		assert.Equal(t, 2, val)

		// 更新值
		lm.Put("apple", 10)
		val, exists = lm.Get("apple")
		assert.True(t, exists)
		assert.Equal(t, 10, val)

		// 检查键存在
		assert.True(t, lm.Contains("cherry"))
		assert.False(t, lm.Contains("grape"))

		// 删除键
		lm.Remove("banana")
		assert.Equal(t, 2, lm.Size())
		assert.False(t, lm.Contains("banana"))
		assert.True(t, lm.Contains("apple"))
		assert.True(t, lm.Contains("cherry"))

		// 清空
		lm.Clear()
		assert.True(t, lm.IsEmpty())
		assert.Equal(t, 0, lm.Size())
	})

	t.Run("插入顺序保持", func(t *testing.T) {
		lm := NewLinkedMap[string, int]()

		// 按顺序插入
		lm.Put("a", 1)
		lm.Put("b", 2)
		lm.Put("c", 3)
		lm.Put("d", 4)

		// 验证键的顺序
		keys := lm.Keys()
		assert.Equal(t, []string{"a", "b", "c", "d"}, keys)

		// 验证值的顺序
		values := lm.Values()
		assert.Equal(t, []int{1, 2, 3, 4}, values)

		// 验证Entries
		entries := lm.Entries()
		assert.Equal(t, 4, len(entries))
		assert.Equal(t, "a", entries[0].K)
		assert.Equal(t, 1, entries[0].V)
		assert.Equal(t, "d", entries[3].K)
		assert.Equal(t, 4, entries[3].V)
	})

	t.Run("更新键移动到尾部", func(t *testing.T) {
		lm := NewLinkedMap[string, int]()

		lm.Put("a", 1)
		lm.Put("b", 2)
		lm.Put("c", 3)

		// 更新中间的键，它会被移到尾部
		lm.Put("b", 20)

		// 验证顺序：a, c, b
		keys := lm.Keys()
		assert.Equal(t, []string{"a", "c", "b"}, keys)

		// 验证值
		val, _ := lm.Get("b")
		assert.Equal(t, 20, val)
	})

	t.Run("First和Last", func(t *testing.T) {
		lm := NewLinkedMap[string, int]()

		lm.Put("a", 1)
		lm.Put("b", 2)
		lm.Put("c", 3)

		// 测试First
		key, val, ok := lm.First()
		assert.True(t, ok)
		assert.Equal(t, "a", key)
		assert.Equal(t, 1, val)

		// 测试Last
		key, val, ok = lm.Last()
		assert.True(t, ok)
		assert.Equal(t, "c", key)
		assert.Equal(t, 3, val)

		// 空映射
		lm2 := NewLinkedMap[string, int]()
		_, _, ok = lm2.First()
		assert.False(t, ok)

		_, _, ok = lm2.Last()
		assert.False(t, ok)
	})

	t.Run("删除头部节点", func(t *testing.T) {
		lm := NewLinkedMap[string, int]()

		lm.Put("a", 1)
		lm.Put("b", 2)
		lm.Put("c", 3)

		// 删除头部
		lm.Remove("a")

		// 验证新头部是b
		key, val, ok := lm.First()
		assert.True(t, ok)
		assert.Equal(t, "b", key)
		assert.Equal(t, 2, val)

		// 验证顺序
		keys := lm.Keys()
		assert.Equal(t, []string{"b", "c"}, keys)
	})

	t.Run("删除尾部节点", func(t *testing.T) {
		lm := NewLinkedMap[string, int]()

		lm.Put("a", 1)
		lm.Put("b", 2)
		lm.Put("c", 3)

		// 删除尾部
		lm.Remove("c")

		// 验证新尾部是b
		key, val, ok := lm.Last()
		assert.True(t, ok)
		assert.Equal(t, "b", key)
		assert.Equal(t, 2, val)

		// 验证顺序
		keys := lm.Keys()
		assert.Equal(t, []string{"a", "b"}, keys)
	})

	t.Run("删除中间节点", func(t *testing.T) {
		lm := NewLinkedMap[string, int]()

		lm.Put("a", 1)
		lm.Put("b", 2)
		lm.Put("c", 3)
		lm.Put("d", 4)

		// 删除中间节点
		lm.Remove("b")

		// 验证顺序
		keys := lm.Keys()
		assert.Equal(t, []string{"a", "c", "d"}, keys)

		// 验证连接正确
		firstKey, _, _ := lm.First()
		assert.Equal(t, "a", firstKey)

		lastKey, _, _ := lm.Last()
		assert.Equal(t, "d", lastKey)
	})

	t.Run("单个节点", func(t *testing.T) {
		lm := NewLinkedMap[string, int]()

		lm.Put("a", 1)

		key, val, ok := lm.First()
		assert.True(t, ok)
		assert.Equal(t, "a", key)
		assert.Equal(t, 1, val)

		key, val, ok = lm.Last()
		assert.True(t, ok)
		assert.Equal(t, "a", key)
		assert.Equal(t, 1, val)

		// 删除唯一节点
		lm.Remove("a")
		assert.True(t, lm.IsEmpty())
	})

	t.Run("整数键", func(t *testing.T) {
		lm := NewLinkedMap[int, string]()

		lm.Put(3, "three")
		lm.Put(1, "one")
		lm.Put(2, "two")

		// 验证顺序保持插入顺序
		keys := lm.Keys()
		assert.Equal(t, []int{3, 1, 2}, keys)
	})
}