package mapx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTreeMap(t *testing.T) {
	t.Run("基本操作", func(t *testing.T) {
		tm := NewTreeMap[string, int]()

		// 测试初始状态
		assert.True(t, tm.IsEmpty())
		assert.Equal(t, 0, tm.Size())

		// 添加键值对
		tm.Put("apple", 1)
		tm.Put("banana", 2)
		tm.Put("cherry", 3)

		assert.False(t, tm.IsEmpty())
		assert.Equal(t, 3, tm.Size())

		// 获取值
		val, exists := tm.Get("apple")
		assert.True(t, exists)
		assert.Equal(t, 1, val)

		val, exists = tm.Get("banana")
		assert.True(t, exists)
		assert.Equal(t, 2, val)

		// 更新值
		tm.Put("apple", 10)
		val, exists = tm.Get("apple")
		assert.True(t, exists)
		assert.Equal(t, 10, val)

		// 检查键存在
		assert.True(t, tm.Contains("cherry"))
		assert.False(t, tm.Contains("grape"))

		// 删除键
		tm.Remove("banana")
		assert.Equal(t, 2, tm.Size())
		assert.False(t, tm.Contains("banana"))
		assert.True(t, tm.Contains("apple"))
		assert.True(t, tm.Contains("cherry"))

		// 清空
		tm.Clear()
		assert.True(t, tm.IsEmpty())
		assert.Equal(t, 0, tm.Size())
	})

	t.Run("有序性", func(t *testing.T) {
		tm := NewTreeMap[int, string]()

		// 乱序插入
		tm.Put(5, "five")
		tm.Put(1, "one")
		tm.Put(3, "three")
		tm.Put(2, "two")
		tm.Put(4, "four")

		// 验证键按排序顺序
		keys := tm.Keys()
		assert.Equal(t, []int{1, 2, 3, 4, 5}, keys)

		// 验证值按排序顺序
		values := tm.Values()
		assert.Equal(t, []string{"one", "two", "three", "four", "five"}, values)

		// 验证Entries
		entries := tm.Entries()
		assert.Equal(t, 5, len(entries))
		for i := 0; i < 5; i++ {
			assert.Equal(t, i+1, entries[i].K)
		}
	})

	t.Run("字符串键排序", func(t *testing.T) {
		tm := NewTreeMap[string, int]()

		tm.Put("cherry", 3)
		tm.Put("apple", 1)
		tm.Put("banana", 2)

		keys := tm.Keys()
		assert.Equal(t, []string{"apple", "banana", "cherry"}, keys)
	})

	t.Run("Min和Max", func(t *testing.T) {
		tm := NewTreeMap[int, string]()

		tm.Put(5, "five")
		tm.Put(1, "one")
		tm.Put(3, "three")

		min, err := tm.Min()
		assert.NoError(t, err)
		assert.Equal(t, 1, min)

		max, err := tm.Max()
		assert.NoError(t, err)
		assert.Equal(t, 5, max)

		// 空树情况
		tm2 := NewTreeMap[int, string]()
		_, err = tm2.Min()
		assert.Error(t, err)

		_, err = tm2.Max()
		assert.Error(t, err)
	})

	t.Run("删除节点", func(t *testing.T) {
		tm := NewTreeMap[int, string]()

		tm.Put(3, "three")
		tm.Put(1, "one")
		tm.Put(4, "four")
		tm.Put(2, "two")

		// 删除叶子节点
		tm.Remove(2)
		assert.False(t, tm.Contains(2))
		assert.Equal(t, 3, tm.Size())

		// 删除中间节点
		tm.Remove(3)
		assert.False(t, tm.Contains(3))
		assert.Equal(t, 2, tm.Size())

		// 验证剩余键有序
		keys := tm.Keys()
		assert.Equal(t, []int{1, 4}, keys)
	})

	t.Run("重复键", func(t *testing.T) {
		tm := NewTreeMap[int, string]()

		tm.Put(1, "one")
		tm.Put(2, "two")
		tm.Put(1, "ONE") // 更新值

		assert.Equal(t, 2, tm.Size())
		val, _ := tm.Get(1)
		assert.Equal(t, "ONE", val)
	})
}