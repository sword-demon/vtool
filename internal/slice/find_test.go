package slice

import (
	"testing"
)

func TestFind(t *testing.T) {
	testCases := []struct {
		name      string
		slice     []int
		predicate func(int) bool
		wantIndex int
		wantVal   int
		wantFound bool
	}{
		{
			name:      "找到第一个大于5的元素",
			slice:     []int{1, 3, 6, 8, 2},
			predicate: func(i int) bool { return i > 5 },
			wantIndex: 2,
			wantVal:   6,
			wantFound: true,
		},
		{
			name:      "找到第一个等于3的元素",
			slice:     []int{1, 2, 3, 4, 3, 5},
			predicate: func(i int) bool { return i == 3 },
			wantIndex: 2,
			wantVal:   3,
			wantFound: true,
		},
		{
			name:      "未找到满足条件的元素",
			slice:     []int{1, 2, 3, 4, 5},
			predicate: func(i int) bool { return i > 10 },
			wantIndex: -1,
			wantVal:   0,
			wantFound: false,
		},
		{
			name:      "空切片查找",
			slice:     []int{},
			predicate: func(i int) bool { return i > 0 },
			wantIndex: -1,
			wantVal:   0,
			wantFound: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			index, val, found := Find(tc.slice, tc.predicate)
			if index != tc.wantIndex {
				t.Errorf("Expected index %d, got %d", tc.wantIndex, index)
			}
			if found != tc.wantFound {
				t.Errorf("Expected found %v, got %v", tc.wantFound, found)
			}
			if found && val != tc.wantVal {
				t.Errorf("Expected value %d, got %d", tc.wantVal, val)
			}
		})
	}
}

func TestFindString(t *testing.T) {
	testCases := []struct {
		name      string
		slice     []string
		predicate func(string) bool
		wantIndex int
		wantVal   string
		wantFound bool
	}{
		{
			name:      "找到字符串匹配",
			slice:     []string{"apple", "banana", "cherry", "date"},
			predicate: func(s string) bool { return s == "cherry" },
			wantIndex: 2,
			wantVal:   "cherry",
			wantFound: true,
		},
		{
			name:      "未找到匹配的字符串",
			slice:     []string{"apple", "banana", "cherry", "date"},
			predicate: func(s string) bool { return s == "grape" },
			wantIndex: -1,
			wantVal:   "",
			wantFound: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			index, val, found := Find(tc.slice, tc.predicate)
			if index != tc.wantIndex {
				t.Errorf("Expected index %d, got %d", tc.wantIndex, index)
			}
			if found != tc.wantFound {
				t.Errorf("Expected found %v, got %v", tc.wantFound, found)
			}
			if found && val != tc.wantVal {
				t.Errorf("Expected value %s, got %s", tc.wantVal, val)
			}
		})
	}
}