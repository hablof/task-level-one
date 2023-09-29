package main

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXxx(t *testing.T) {
	tests := []struct {
		name    string
		arr     []int
		idx     int
		want    []int
		wantErr error
	}{
		{
			name:    "1",
			arr:     []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			idx:     0,
			want:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			wantErr: nil,
		},
		{
			name:    "2",
			arr:     []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			idx:     5,
			want:    []int{0, 1, 2, 3, 4, 6, 7, 8, 9, 10},
			wantErr: nil,
		},
		{
			name:    "3",
			arr:     []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			idx:     8,
			want:    []int{0, 1, 2, 3, 4, 5, 6, 7, 9, 10},
			wantErr: nil,
		},
		{
			name:    "4",
			arr:     []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			idx:     10,
			want:    []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			wantErr: nil,
		},
		{
			name:    "5",
			arr:     []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			idx:     11,
			want:    []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			wantErr: ErrNotInBounds,
		},
		{
			name:    "5",
			arr:     []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			idx:     -1,
			want:    []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			wantErr: ErrNotInBounds,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cloneArr := slices.Clone(tt.arr)

			err1 := DeleteIthFromSliceWithSavedOrder(&tt.arr, tt.idx)
			err2 := DeleteIthFromSliceNoOrder(&cloneArr, tt.idx)

			assert.Equal(t, tt.wantErr, err1)
			assert.Equal(t, tt.wantErr, err2)
			assert.Equal(t, tt.want, tt.arr)
			assert.NotContains(t, cloneArr, tt.idx) // не очень правильно, но логично
		})
	}
}
