package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSkipList(t *testing.T) {
	t.Run("基本操作", func(t *testing.T) {
		s := NewSkipList[int]()

		// 测试初始状态
		assert.True(t, s.IsEmpty())
		assert.Equal(t, 0, s.Size())

		// 插入元素
		s.Insert(5)
		s.Insert(2)
		s.Insert(8)
		s.Insert(1)
		s.Insert(9)

		assert.False(t, s.IsEmpty())
		assert.Equal(t, 5, s.Size())

		// 查找元素
		assert.True(t, s.Contains(5))
		assert.True(t, s.Contains(2))
		assert.True(t, s.Contains(8))
		assert.True(t, s.Contains(1))
		assert.True(t, s.Contains(9))
		assert.False(t, s.Contains(10))

		// 测试Min和Max
		min, err := s.Min()
		assert.NoError(t, err)
		assert.Equal(t, 1, min)

		max, err := s.Max()
		assert.NoError(t, err)
		assert.Equal(t, 9, max)

		// 检查有序性
		slice := s.ToSlice()
		assert.Equal(t, 5, len(slice))
		// 跳表保证有序性
		for i := 0; i < len(slice)-1; i++ {
			assert.True(t, slice[i] <= slice[i+1])
		}
	})

	t.Run("删除操作", func(t *testing.T) {
		s := NewSkipList[int]()

		// 插入测试数据
		values := []int{5, 2, 8, 1, 9, 3, 7, 4, 6}
		for _, v := range values {
			s.Insert(v)
		}

		// 删除中间元素
		assert.True(t, s.Remove(5))
		assert.Equal(t, 8, s.Size())
		assert.False(t, s.Contains(5))

		// 删除头部元素
		assert.True(t, s.Remove(1))
		assert.Equal(t, 7, s.Size())
		assert.False(t, s.Contains(1))

		// 删除尾部元素
		assert.True(t, s.Remove(9))
		assert.Equal(t, 6, s.Size())
		assert.False(t, s.Contains(9))

		// 删除不存在的元素
		assert.False(t, s.Remove(100))

		// 清空
		s.Clear()
		assert.True(t, s.IsEmpty())
		assert.Equal(t, 0, s.Size())
	})

	t.Run("去重", func(t *testing.T) {
		s := NewSkipList[int]()

		// 插入重复元素
		s.Insert(1)
		s.Insert(2)
		s.Insert(1) // 重复，不应该插入
		s.Insert(3)
		s.Insert(2) // 重复，不应该插入

		assert.Equal(t, 3, s.Size())
		slice := s.ToSlice()
		assert.Equal(t, []int{1, 2, 3}, slice)
	})

	t.Run("字符串跳表", func(t *testing.T) {
		s := NewSkipList[string]()

		s.Insert("banana")
		s.Insert("apple")
		s.Insert("cherry")

		assert.Equal(t, 3, s.Size())
		assert.True(t, s.Contains("banana"))
		assert.False(t, s.Contains("grape"))

		slice := s.ToSlice()
		assert.Equal(t, []string{"apple", "banana", "cherry"}, slice)

		min, err := s.Min()
		assert.NoError(t, err)
		assert.Equal(t, "apple", min)

		max, err := s.Max()
		assert.NoError(t, err)
		assert.Equal(t, "cherry", max)
	})

	t.Run("大随机数据", func(t *testing.T) {
		s := NewSkipList[int]()

		// 插入100个随机数
		for i := 0; i < 100; i++ {
			s.Insert(i)
		}

		assert.Equal(t, 100, s.Size())

		// 随机删除一些元素
		for i := 0; i < 50; i++ {
			s.Remove(i)
		}

		assert.Equal(t, 50, s.Size())

		// 检查是否还有重复（应该没有）
		slice := s.ToSlice()
		seen := make(map[int]bool)
		for _, v := range slice {
			assert.False(t, seen[v])
			seen[v] = true
		}
	})
}