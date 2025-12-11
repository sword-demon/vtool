package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	t.Run("基本操作", func(t *testing.T) {
		q := NewQueue[int]()

		// 测试初始状态
		assert.True(t, q.IsEmpty())
		assert.Equal(t, 0, q.Size())

		// 入队
		q.Enqueue(1)
		q.Enqueue(2)
		q.Enqueue(3)

		assert.False(t, q.IsEmpty())
		assert.Equal(t, 3, q.Size())

		// 查看队首
		val, err := q.Peek()
		assert.NoError(t, err)
		assert.Equal(t, 1, val)

		// 出队
		val, err = q.Dequeue()
		assert.NoError(t, err)
		assert.Equal(t, 1, val)

		val, err = q.Dequeue()
		assert.NoError(t, err)
		assert.Equal(t, 2, val)

		// 再次查看
		val, err = q.Peek()
		assert.NoError(t, err)
		assert.Equal(t, 3, val)

		// 最后出队
		val, err = q.Dequeue()
		assert.NoError(t, err)
		assert.Equal(t, 3, val)

		// 队列为空
		assert.True(t, q.IsEmpty())
		assert.Equal(t, 0, q.Size())
	})

	t.Run("错误情况", func(t *testing.T) {
		q := NewQueue[int]()

		// 空队列出队
		_, err := q.Dequeue()
		assert.Error(t, err)

		// 空队列查看
		_, err = q.Peek()
		assert.Error(t, err)
	})

	t.Run("清空队列", func(t *testing.T) {
		q := NewQueue[int]()

		q.Enqueue(1)
		q.Enqueue(2)
		q.Enqueue(3)

		assert.Equal(t, 3, q.Size())

		q.Clear()

		assert.True(t, q.IsEmpty())
		assert.Equal(t, 0, q.Size())
	})

	t.Run("ToSlice", func(t *testing.T) {
		q := NewQueue[int]()

		q.Enqueue(1)
		q.Enqueue(2)
		q.Enqueue(3)

		slice := q.ToSlice()
		assert.Equal(t, []int{1, 2, 3}, slice)
	})

	t.Run("字符串队列", func(t *testing.T) {
		q := NewQueue[string]()

		q.Enqueue("apple")
		q.Enqueue("banana")
		q.Enqueue("cherry")

		assert.Equal(t, 3, q.Size())

		val, err := q.Dequeue()
		assert.NoError(t, err)
		assert.Equal(t, "apple", val)

		val, err = q.Peek()
		assert.NoError(t, err)
		assert.Equal(t, "banana", val)
	})
}