package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPriorityQueue(t *testing.T) {
	t.Run("基本操作", func(t *testing.T) {
		pq := NewPriorityQueue[int]()

		// 测试初始状态
		assert.True(t, pq.IsEmpty())
		assert.Equal(t, 0, pq.Size())

		// 入队不同优先级的元素
		pq.Enqueue(3, 3) // 优先级3
		pq.Enqueue(1, 1) // 优先级1（最高）
		pq.Enqueue(2, 2) // 优先级2

		assert.False(t, pq.IsEmpty())
		assert.Equal(t, 3, pq.Size())

		// 查看最高优先级元素
		val, priority, err := pq.Peek()
		assert.NoError(t, err)
		assert.Equal(t, 1, val)
		assert.Equal(t, 1, priority)

		// 按优先级出队
		val, priority, err = pq.Dequeue()
		assert.NoError(t, err)
		assert.Equal(t, 1, val)
		assert.Equal(t, 1, priority)

		val, priority, err = pq.Dequeue()
		assert.NoError(t, err)
		assert.Equal(t, 2, val)
		assert.Equal(t, 2, priority)

		val, priority, err = pq.Dequeue()
		assert.NoError(t, err)
		assert.Equal(t, 3, val)
		assert.Equal(t, 3, priority)

		// 队列为空
		assert.True(t, pq.IsEmpty())
		assert.Equal(t, 0, pq.Size())
	})

	t.Run("相同优先级", func(t *testing.T) {
		pq := NewPriorityQueue[int]()

		// 插入相同优先级的元素
		pq.Enqueue(1, 1)
		pq.Enqueue(2, 1)
		pq.Enqueue(3, 1)

		// 优先级相同时，出队顺序可能不是严格的FIFO
		// 但所有元素都应该被正确出队
		values := make([]int, 0, 3)
		for i := 0; i < 3; i++ {
			val, _, err := pq.Dequeue()
			assert.NoError(t, err)
			values = append(values, val)
		}

		// 验证所有元素都被出队
		assert.Equal(t, 3, len(values))
		assert.Contains(t, values, 1)
		assert.Contains(t, values, 2)
		assert.Contains(t, values, 3)
	})

	t.Run("错误情况", func(t *testing.T) {
		pq := NewPriorityQueue[int]()

		// 空队列出队
		_, _, err := pq.Dequeue()
		assert.Error(t, err)

		// 空队列查看
		_, _, err = pq.Peek()
		assert.Error(t, err)
	})

	t.Run("清空队列", func(t *testing.T) {
		pq := NewPriorityQueue[int]()

		pq.Enqueue(1, 1)
		pq.Enqueue(2, 2)
		pq.Enqueue(3, 3)

		assert.Equal(t, 3, pq.Size())

		pq.Clear()

		assert.True(t, pq.IsEmpty())
		assert.Equal(t, 0, pq.Size())
	})

	t.Run("ToSlice", func(t *testing.T) {
		pq := NewPriorityQueue[int]()

		pq.Enqueue(3, 3)
		pq.Enqueue(1, 1)
		pq.Enqueue(2, 2)

		slice := pq.ToSlice()
		assert.Equal(t, 3, len(slice))

		// 检查第一个元素是优先级最高的
		assert.Equal(t, 1, slice[0].Value)
		assert.Equal(t, 1, slice[0].Priority)
	})

	t.Run("字符串优先级队列", func(t *testing.T) {
		pq := NewPriorityQueue[string]()

		pq.Enqueue("low", 3)
		pq.Enqueue("high", 1)
		pq.Enqueue("medium", 2)

		val, priority, err := pq.Dequeue()
		assert.NoError(t, err)
		assert.Equal(t, "high", val)
		assert.Equal(t, 1, priority)

		val, priority, err = pq.Peek()
		assert.NoError(t, err)
		assert.Equal(t, "medium", val)
		assert.Equal(t, 2, priority)
	})

	t.Run("复杂场景", func(t *testing.T) {
		pq := NewPriorityQueue[string]()

		// 模拟任务调度
		pq.Enqueue("task1", 3)
		pq.Enqueue("urgent", 1)
		pq.Enqueue("task2", 2)
		pq.Enqueue("normal", 4)

		// 优先处理紧急任务
		val, priority, err := pq.Dequeue()
		assert.NoError(t, err)
		assert.Equal(t, "urgent", val)
		assert.Equal(t, 1, priority)

		// 再处理其他任务
		val, _, err = pq.Dequeue()
		assert.NoError(t, err)
		assert.Equal(t, "task2", val)

		val, _, err = pq.Dequeue()
		assert.NoError(t, err)
		assert.Equal(t, "task1", val)

		val, _, err = pq.Dequeue()
		assert.NoError(t, err)
		assert.Equal(t, "normal", val)
	})
}