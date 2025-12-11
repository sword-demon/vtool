package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	testCases := []struct {
		name     string
		src      []int
		mapper   func(int) int
		expected []int
	}{
		{
			name:     "每个元素乘以2",
			src:      []int{1, 2, 3, 4},
			mapper:   func(i int) int { return i * 2 },
			expected: []int{2, 4, 6, 8},
		},
		{
			name:     "整数加10",
			src:      []int{1, 2, 3},
			mapper:   func(i int) int { return i + 10 },
			expected: []int{11, 12, 13},
		},
		{
			name:     "空切片",
			src:      []int{},
			mapper:   func(i int) int { return i * 2 },
			expected: []int{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Map(tc.src, tc.mapper)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestReduce(t *testing.T) {
	testCases := []struct {
		name     string
		src      []int
		reducer  func(int, int) int
		initial  int
		expected int
	}{
		{
			name:     "求和",
			src:      []int{1, 2, 3, 4, 5},
			reducer:  func(acc, val int) int { return acc + val },
			initial:  0,
			expected: 15,
		},
		{
			name:     "求积",
			src:      []int{2, 3, 4},
			reducer:  func(acc, val int) int { return acc * val },
			initial:  1,
			expected: 24,
		},
		{
			name: "找最大值",
			src:  []int{3, 7, 2, 9, 1},
			reducer: func(acc, val int) int {
				if val > acc {
					return val
				}
				return acc
			},
			initial:  0,
			expected: 9,
		},
		{
			name:     "初始值不为0的求和",
			src:      []int{1, 2, 3},
			reducer:  func(acc, val int) int { return acc + val },
			initial:  10,
			expected: 16,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Reduce(tc.src, tc.reducer, tc.initial)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestFilter(t *testing.T) {
	testCases := []struct {
		name      string
		src       []int
		predicate func(int) bool
		expected  []int
	}{
		{
			name:      "过滤出偶数",
			src:       []int{1, 2, 3, 4, 5, 6},
			predicate: func(i int) bool { return i%2 == 0 },
			expected:  []int{2, 4, 6},
		},
		{
			name:      "过滤出大于5的数",
			src:       []int{1, 6, 3, 8, 2, 9},
			predicate: func(i int) bool { return i > 5 },
			expected:  []int{6, 8, 9},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Filter(tc.src, tc.predicate)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestFilterString(t *testing.T) {
	testCases := []struct {
		name      string
		src       []string
		predicate func(string) bool
		expected  []string
	}{
		{
			name:      "过滤出字符串长度大于3的",
			src:       []string{"a", "ab", "abc", "abcd", "abcde"},
			predicate: func(s string) bool { return len(s) > 3 },
			expected:  []string{"abcd", "abcde"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Filter(tc.src, tc.predicate)
			assert.Equal(t, tc.expected, result)
		})
	}
}
