package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortedSet(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "基本排序和去重",
			input:    []int{3, 1, 2, 3, 1, 2},
			expected: []int{1, 2, 3},
		},
		{
			name:     "空切片",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "已有序的切片",
			input:    []int{1, 2, 2, 3, 3, 3},
			expected: []int{1, 2, 3},
		},
		{
			name:     "降序切片",
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "包含重复元素",
			input:    []int{1, 1, 1, 1},
			expected: []int{1},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := SortedSet(tc.input)
			assert.Equal(t, tc.expected, result)

			// 确保原切片未被修改
			if len(tc.input) > 0 {
				assert.NotEqual(t, result, tc.input)
			}
		})
	}
}

func TestSortedSetDesc(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "基本降序排序和去重",
			input:    []int{3, 1, 2, 3, 1, 2},
			expected: []int{3, 2, 1},
		},
		{
			name:     "空切片",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "升序切片",
			input:    []int{1, 2, 3, 3, 4},
			expected: []int{4, 3, 2, 1},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := SortedSetDesc(tc.input)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestUnique(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "基本去重",
			input:    []int{3, 1, 2, 3, 1, 2},
			expected: []int{3, 1, 2},
		},
		{
			name:     "空切片",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "保持原有顺序",
			input:    []int{1, 2, 3, 2, 1, 4},
			expected: []int{1, 2, 3, 4},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Unique(tc.input)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestSortedSetString(t *testing.T) {
	input := []string{"cherry", "apple", "banana", "apple", "cherry"}
	expected := []string{"apple", "banana", "cherry"}

	result := SortedSet(input)
	assert.Equal(t, expected, result)
}
