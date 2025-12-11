package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnion(t *testing.T) {
	testCases := []struct {
		name     string
		src1     []int
		src2     []int
		expected []int
	}{
		{
			name:     "两个无重复元素的切片求并集",
			src1:     []int{1, 2, 3},
			src2:     []int{4, 5, 6},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "有重复元素的切片求并集",
			src1:     []int{1, 2, 3},
			src2:     []int{2, 3, 4},
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "第一个切片为空",
			src1:     []int{},
			src2:     []int{1, 2, 3},
			expected: []int{1, 2, 3},
		},
		{
			name:     "第二个切片为空",
			src1:     []int{1, 2, 3},
			src2:     []int{},
			expected: []int{1, 2, 3},
		},
		{
			name:     "两个空切片",
			src1:     []int{},
			src2:     []int{},
			expected: []int{},
		},
		{
			name:     "完全重复的切片",
			src1:     []int{1, 2, 3},
			src2:     []int{1, 2, 3},
			expected: []int{1, 2, 3},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Union(tc.src1, tc.src2)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestUnionString(t *testing.T) {
	testCases := []struct {
		name     string
		src1     []string
		src2     []string
		expected []string
	}{
		{
			name:     "字符串切片求并集",
			src1:     []string{"apple", "banana"},
			src2:     []string{"cherry", "date", "apple"},
			expected: []string{"apple", "banana", "cherry", "date"},
		},
		{
			name:     "字符串切片完全重复",
			src1:     []string{"apple", "banana"},
			src2:     []string{"apple", "banana"},
			expected: []string{"apple", "banana"},
		},
		{
			name:     "字符串切片一个为空",
			src1:     []string{},
			src2:     []string{"apple", "banana"},
			expected: []string{"apple", "banana"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Union(tc.src1, tc.src2)
			assert.Equal(t, tc.expected, result)
		})
	}
}