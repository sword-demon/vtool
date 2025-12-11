package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashSet(t *testing.T) {
	t.Run("基本操作", func(t *testing.T) {
		s := NewHashSet[int]()

		// 测试初始状态
		assert.True(t, s.IsEmpty())
		assert.Equal(t, 0, s.Size())

		// 添加元素
		s.Add(1)
		s.Add(2)
		s.Add(3)

		assert.False(t, s.IsEmpty())
		assert.Equal(t, 3, s.Size())
		assert.True(t, s.Contains(1))
		assert.True(t, s.Contains(2))
		assert.True(t, s.Contains(3))
		assert.False(t, s.Contains(4))

		// 测试去重
		s.Add(1)
		assert.Equal(t, 3, s.Size())

		// 删除元素
		s.Remove(2)
		assert.Equal(t, 2, s.Size())
		assert.False(t, s.Contains(2))
		assert.True(t, s.Contains(1))
		assert.True(t, s.Contains(3))

		// 清空集合
		s.Clear()
		assert.True(t, s.IsEmpty())
		assert.Equal(t, 0, s.Size())
	})

	t.Run("字符串集合", func(t *testing.T) {
		s := NewHashSet[string]()

		s.Add("apple")
		s.Add("banana")
		s.Add("cherry")

		assert.Equal(t, 3, s.Size())
		assert.True(t, s.Contains("banana"))
		assert.False(t, s.Contains("grape"))
	})

	t.Run("ToSlice", func(t *testing.T) {
		s := NewHashSet[int]()
		items := []int{3, 1, 2, 4, 5}

		for _, item := range items {
			s.Add(item)
		}

		slice := s.ToSlice()
		assert.Len(t, slice, 5)

		// 检查所有元素都存在
		for _, item := range items {
			assert.Contains(t, slice, item)
		}
	})
}

func TestHashSetUnion(t *testing.T) {
	s1 := NewHashSet[int]()
	s1.Add(1)
	s1.Add(2)
	s1.Add(3)

	s2 := NewHashSet[int]()
	s2.Add(3)
	s2.Add(4)
	s2.Add(5)

	union := s1.Union(s2)

	assert.Equal(t, 5, union.Size())
	assert.True(t, union.Contains(1))
	assert.True(t, union.Contains(2))
	assert.True(t, union.Contains(3))
	assert.True(t, union.Contains(4))
	assert.True(t, union.Contains(5))
}

func TestHashSetIntersect(t *testing.T) {
	s1 := NewHashSet[int]()
	s1.Add(1)
	s1.Add(2)
	s1.Add(3)

	s2 := NewHashSet[int]()
	s2.Add(2)
	s2.Add(3)
	s2.Add(4)

	intersect := s1.Intersect(s2)

	assert.Equal(t, 2, intersect.Size())
	assert.True(t, intersect.Contains(2))
	assert.True(t, intersect.Contains(3))
	assert.False(t, intersect.Contains(1))
	assert.False(t, intersect.Contains(4))
}

func TestHashSetDifference(t *testing.T) {
	s1 := NewHashSet[int]()
	s1.Add(1)
	s1.Add(2)
	s1.Add(3)

	s2 := NewHashSet[int]()
	s2.Add(2)
	s2.Add(3)
	s2.Add(4)

	diff := s1.Difference(s2)

	assert.Equal(t, 1, diff.Size())
	assert.True(t, diff.Contains(1))
	assert.False(t, diff.Contains(2))
	assert.False(t, diff.Contains(3))
	assert.False(t, diff.Contains(4))
}
