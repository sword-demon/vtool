package maps

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	testCases := []struct {
		name     string
		src      map[string]int
		key      string
		expected int
		exists   bool
	}{
		{
			name:     "key存在",
			src:      map[string]int{"apple": 1, "banana": 2, "cherry": 3},
			key:      "banana",
			expected: 2,
			exists:   true,
		},
		{
			name:     "key不存在",
			src:      map[string]int{"apple": 1, "banana": 2},
			key:      "grape",
			expected: 0,
			exists:   false,
		},
		{
			name:     "空map",
			src:      map[string]int{},
			key:      "any",
			expected: 0,
			exists:   false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			val, exists := Get(tc.src, tc.key)
			assert.Equal(t, tc.expected, val)
			assert.Equal(t, tc.exists, exists)
		})
	}
}

func TestHas(t *testing.T) {
	testCases := []struct {
		name     string
		src      map[string]int
		key      string
		expected bool
	}{
		{
			name:     "key存在",
			src:      map[string]int{"apple": 1, "banana": 2},
			key:      "banana",
			expected: true,
		},
		{
			name:     "key不存在",
			src:      map[string]int{"apple": 1, "banana": 2},
			key:      "grape",
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Has(tc.src, tc.key)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestKeys(t *testing.T) {
	testCases := []struct {
		name     string
		src      map[string]int
		expected []string
	}{
		{
			name:     "多个key",
			src:      map[string]int{"apple": 1, "banana": 2, "cherry": 3},
			expected: []string{"apple", "banana", "cherry"},
		},
		{
			name:     "空map",
			src:      map[string]int{},
			expected: []string{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Keys(tc.src)
			// 由于map的迭代顺序不确定，我们检查长度和内容
			assert.Len(t, result, len(tc.expected))
			for _, key := range tc.expected {
				assert.Contains(t, result, key)
			}
		})
	}
}

func TestValues(t *testing.T) {
	testCases := []struct {
		name     string
		src      map[string]int
		expected []int
	}{
		{
			name:     "多个value",
			src:      map[string]int{"apple": 1, "banana": 2, "cherry": 3},
			expected: []int{1, 2, 3},
		},
		{
			name:     "空map",
			src:      map[string]int{},
			expected: []int{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Values(tc.src)
			// 由于map的迭代顺序不确定，我们检查长度和内容
			assert.Len(t, result, len(tc.expected))
			for _, val := range tc.expected {
				assert.Contains(t, result, val)
			}
		})
	}
}
