package slice

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDelete(t *testing.T) {
	testCases := []struct {
		name      string
		slice     []int
		index     int
		wantSlice []int
		wantErr   error
	}{
		{
			name:      "删除第一个元素",
			slice:     []int{1, 2, 3, 4},
			index:     0,
			wantSlice: []int{2, 3, 4},
		},
		{
			name:      "删除中间元素",
			slice:     []int{1, 2, 3, 4},
			index:     1,
			wantSlice: []int{1, 3, 4},
		},
		{
			name:      "删除最后一个元素",
			slice:     []int{1, 2, 3, 4},
			index:     3,
			wantSlice: []int{1, 2, 3},
		},
		{
			name:      "删除单个元素的切片",
			slice:     []int{1},
			index:     0,
			wantSlice: []int{},
		},
		{
			name:      "错误情况 - index < 0",
			slice:     []int{1, 2, 3},
			index:     -1,
			wantSlice: nil,
			wantErr:   errors.New("index is out of range"),
		},
		{
			name:      "错误情况 - index >= len",
			slice:     []int{1, 2, 3},
			index:     3,
			wantSlice: nil,
			wantErr:   errors.New("index is out of range"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := Delete(tc.slice, tc.index)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantSlice, res)
		})
	}
}
