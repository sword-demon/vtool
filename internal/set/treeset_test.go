package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTreeSet(t *testing.T) {
	t.Run("基本操作", func(t *testing.T) {
		s := NewTreeSet[int]()

		// 测试初始状态
		assert.True(t, s.IsEmpty())
		assert.Equal(t, 0, s.Size())

		// 添加元素
		s.Add(3)
		s.Add(1)
		s.Add(2)

		assert.False(t, s.IsEmpty())
		assert.Equal(t, 3, s.Size())
		assert.True(t, s.Contains(1))
		assert.True(t, s.Contains(2))
		assert.True(t, s.Contains(3))

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

	t.Run("有序性", func(t *testing.T) {
		s := NewTreeSet[int]()

		// 添加乱序元素
		s.Add(5)
		s.Add(1)
		s.Add(3)
		s.Add(2)
		s.Add(4)

		// 检查是否有序
		slice := s.ToSlice()
		assert.Equal(t, []int{1, 2, 3, 4, 5}, slice)
	})

	t.Run("字符串集合", func(t *testing.T) {
		s := NewTreeSet[string]()

		s.Add("cherry")
		s.Add("apple")
		s.Add("banana")

		assert.Equal(t, 3, s.Size())
		assert.True(t, s.Contains("banana"))
		assert.False(t, s.Contains("grape"))

		// 检查字符串排序
		slice := s.ToSlice()
		assert.Equal(t, []string{"apple", "banana", "cherry"}, slice)
	})

	t.Run("ToSlice", func(t *testing.T) {
		s := NewTreeSet[int]()
		items := []int{3, 1, 2, 4, 5}

		for _, item := range items {
			s.Add(item)
		}

		slice := s.ToSlice()
		assert.Equal(t, []int{1, 2, 3, 4, 5}, slice)
	})
}

func TestTreeSetUnion(t *testing.T) {
	s1 := NewTreeSet[int]()
	s1.Add(1)
	s1.Add(2)
	s1.Add(3)

	s2 := NewTreeSet[int]()
	s2.Add(3)
	s2.Add(4)
	s2.Add(5)

	union := s1.Union(s2)

	assert.Equal(t, 5, union.Size())
	// 检查并集是否有序
	assert.Equal(t, []int{1, 2, 3, 4, 5}, union.ToSlice())
}

func TestTreeSetIntersect(t *testing.T) {
	s1 := NewTreeSet[int]()
	s1.Add(1)
	s1.Add(2)
	s1.Add(3)

	s2 := NewTreeSet[int]()
	s2.Add(2)
	s2.Add(3)
	s2.Add(4)

	intersect := s1.Intersect(s2)

	assert.Equal(t, 2, intersect.Size())
	assert.Equal(t, []int{2, 3}, intersect.ToSlice())
}

func TestTreeSetDifference(t *testing.T) {
	s1 := NewTreeSet[int]()
	s1.Add(1)
	s1.Add(2)
	s1.Add(3)

	s2 := NewTreeSet[int]()
	s2.Add(2)
	s2.Add(3)
	s2.Add(4)

	diff := s1.Difference(s2)

	assert.Equal(t, 1, diff.Size())
	assert.Equal(t, []int{1}, diff.ToSlice())
}
