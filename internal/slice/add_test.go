package slice

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	testCases := []struct {
		name      string
		slice     []int
		addVal    int
		index     int
		wantSlice []int
		wantErr   error
	}{
		{
			name:      "index 0 - 在头部添加",
			slice:     []int{123, 100},
			addVal:    233,
			index:     0,
			wantSlice: []int{233, 123, 100},
		},
		{
			name:      "在中间位置添加",
			slice:     []int{1, 2, 3, 4},
			addVal:    99,
			index:     2,
			wantSlice: []int{1, 2, 99, 3, 4},
		},
		{
			name:      "在末尾添加 - index == len",
			slice:     []int{1, 2, 3},
			addVal:    4,
			index:     3,
			wantSlice: []int{1, 2, 3, 4},
		},
		{
			name:      "空切片添加",
			slice:     []int{},
			addVal:    1,
			index:     0,
			wantSlice: []int{1},
		},
		{
			name:      "错误情况 - index < 0",
			slice:     []int{1, 2, 3},
			addVal:    99,
			index:     -1,
			wantSlice: nil,
			wantErr:   errors.New("index is out of range"),
		},
		{
			name:      "错误情况 - index > len",
			slice:     []int{1, 2, 3},
			addVal:    99,
			index:     4,
			wantSlice: nil,
			wantErr:   errors.New("index is out of range"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := Add(tc.slice, tc.addVal, tc.index)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantSlice, res)
		})
	}
}
