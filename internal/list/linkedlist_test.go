package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedList(t *testing.T) {
	t.Run("基本操作", func(t *testing.T) {
		l := NewLinkedList[int]()

		// 测试初始状态
		assert.True(t, l.IsEmpty())
		assert.Equal(t, 0, l.Size())

		// 添加元素
		l.Add(1)
		l.Add(2)
		l.Add(3)

		assert.False(t, l.IsEmpty())
		assert.Equal(t, 3, l.Size())

		// 获取元素
		val, err := l.Get(0)
		assert.NoError(t, err)
		assert.Equal(t, 1, val)

		val, err = l.Get(1)
		assert.NoError(t, err)
		assert.Equal(t, 2, val)

		val, err = l.Get(2)
		assert.NoError(t, err)
		assert.Equal(t, 3, val)

		// 设置元素
		err = l.Set(1, 99)
		assert.NoError(t, err)

		val, err = l.Get(1)
		assert.NoError(t, err)
		assert.Equal(t, 99, val)

		// 查找元素
		assert.True(t, l.Contains(99))
		assert.False(t, l.Contains(2))
		assert.Equal(t, 0, l.IndexOf(1))
		assert.Equal(t, 1, l.IndexOf(99))
		assert.Equal(t, 2, l.IndexOf(3))
		assert.Equal(t, -1, l.IndexOf(100))
	})

	t.Run("插入操作", func(t *testing.T) {
		l := NewLinkedList[int]()

		// 插入到头部
		l.Insert(0, 2)
		l.Insert(0, 1)
		assert.Equal(t, []int{1, 2}, l.ToSlice())

		// 插入到中间
		l.Insert(1, 99)
		assert.Equal(t, []int{1, 99, 2}, l.ToSlice())

		// 插入到尾部
		l.Insert(3, 3)
		assert.Equal(t, []int{1, 99, 2, 3}, l.ToSlice())

		// 错误情况
		err := l.Insert(-1, 100)
		assert.Error(t, err)

		err = l.Insert(10, 100)
		assert.Error(t, err)
	})

	t.Run("删除操作", func(t *testing.T) {
		l := NewLinkedList[int]()
		l.FromSlice([]int{1, 2, 3, 4, 5})

		// 删除头部
		err := l.Remove(0)
		assert.NoError(t, err)
		assert.Equal(t, []int{2, 3, 4, 5}, l.ToSlice())

		// 删除尾部
		err = l.Remove(2)
		assert.NoError(t, err)
		assert.Equal(t, []int{2, 3, 5}, l.ToSlice())

		// 删除中间元素
		err = l.Remove(1)
		assert.NoError(t, err)
		assert.Equal(t, []int{2, 5}, l.ToSlice())

		// 清空
		l.Clear()
		assert.True(t, l.IsEmpty())
		assert.Equal(t, 0, l.Size())

		// 错误情况
		err = l.Remove(0)
		assert.Error(t, err)

		l.Add(1)
		err = l.Remove(1)
		assert.Error(t, err)
	})

	t.Run("字符串链表", func(t *testing.T) {
		l := NewLinkedList[string]()

		l.Add("apple")
		l.Add("banana")
		l.Add("cherry")

		assert.Equal(t, 3, l.Size())
		assert.True(t, l.Contains("banana"))
		assert.False(t, l.Contains("grape"))

		val, err := l.Get(1)
		assert.NoError(t, err)
		assert.Equal(t, "banana", val)
	})

	t.Run("FromSlice", func(t *testing.T) {
		l := NewLinkedList[int]()
		values := []int{5, 3, 8, 1, 9}

		l.FromSlice(values)
		assert.Equal(t, 5, l.Size())

		slice := l.ToSlice()
		assert.Equal(t, values, slice)
	})
}
