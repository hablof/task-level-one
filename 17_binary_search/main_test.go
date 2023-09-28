package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_binarySearch(t *testing.T) {

	t.Run("integers", func(t *testing.T) {
		tests := []struct {
			name    string
			nums    []int
			target  int
			want    int
			wantErr error
		}{
			{
				name:    "not found",
				nums:    []int{3, 5, 11, 17, 31, 41, 59, 67, 83, 109, 127, 157},
				target:  0,
				want:    -1,
				wantErr: ErrNotFound,
			},
			{
				name:    "left edge",
				nums:    []int{3, 5, 11, 17, 31, 41, 59, 67, 83, 109, 127, 157},
				target:  3,
				want:    0,
				wantErr: nil,
			},
			{
				name:    "right edge",
				nums:    []int{3, 5, 11, 17, 31, 41, 59, 67, 83, 109, 127, 157},
				target:  157,
				want:    11,
				wantErr: nil,
			},
			{
				name:    "one more",
				nums:    []int{-1, 0, 3, 5, 9, 12},
				target:  9,
				want:    4,
				wantErr: nil,
			},
			{
				name:    "not found",
				nums:    []int{-1, 0, 3, 5, 9, 12},
				target:  2,
				want:    -1,
				wantErr: ErrNotFound,
			},
			{
				name:    "empty slice",
				nums:    []int{},
				target:  2,
				want:    -1,
				wantErr: ErrNotFound,
			},
			{
				name:    "one value slice",
				nums:    []int{1},
				target:  1,
				want:    0,
				wantErr: nil,
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := binarySearch(tt.nums, tt.target)
				assert.Equal(t, tt.want, got)
				assert.Equal(t, tt.wantErr, err)
			})
		}
	})

	t.Run("floats", func(t *testing.T) {
		tests := []struct {
			name    string
			floats  []float64
			target  float64
			want    int
			wantErr error
		}{
			{
				name:    "not found",
				floats:  []float64{3.7, 5, 11, 17, 31, 41.15, 59, 67, 83, 109.45, 127, 157.127},
				target:  43.5,
				want:    -1,
				wantErr: ErrNotFound,
			},
			{
				name:    "left edge",
				floats:  []float64{3.7, 5, 11, 17, 31, 41.15, 59, 67, 83, 109.45, 127, 157.127},
				target:  3.7,
				want:    0,
				wantErr: nil,
			},
			{
				name:    "right edge",
				floats:  []float64{33.7, 5, 11, 17, 31, 41.15, 59, 67, 83, 109.45, 127, 157.127},
				target:  157.127,
				want:    11,
				wantErr: nil,
			},
			{
				name:    "one more",
				floats:  []float64{-1, 0, 3, 5, 9, 12},
				target:  9,
				want:    4,
				wantErr: nil,
			},
			{
				name:    "not found",
				floats:  []float64{-1, 0, 3, 5, 9, 12},
				target:  9.1,
				want:    -1,
				wantErr: ErrNotFound,
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := binarySearch(tt.floats, tt.target)
				assert.Equal(t, tt.want, got)
				assert.Equal(t, tt.wantErr, err)
			})
		}
	})

	t.Run("strings", func(t *testing.T) {
		tests := []struct {
			name    string
			strings []string
			target  string
			want    int
			wantErr error
		}{
			{
				name:    "not found",
				strings: []string{"Кириллица", "Применяется", "Программа", "алфавитном", "в", "и", "компоненты", "латиница", "на", "порядке", "разбором", "слова", "сортировке", "учебе"},
				target:  "flex",
				want:    -1,
				wantErr: ErrNotFound,
			},
			{
				name:    "left edge",
				strings: []string{"Кириллица", "Применяется", "Программа", "алфавитном", "в", "и", "компоненты", "латиница", "на", "порядке", "разбором", "слова", "сортировке", "учебе"},
				target:  "Кириллица",
				want:    0,
				wantErr: nil,
			},
			{
				name:    "right edge",
				strings: []string{"Кириллица", "Применяется", "Программа", "алфавитном", "в", "и", "компоненты", "латиница", "на", "порядке", "разбором", "слова", "сортировке", "учебе"},
				target:  "учебе",
				want:    13,
				wantErr: nil,
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := binarySearch(tt.strings, tt.target)
				assert.Equal(t, tt.want, got)
				assert.Equal(t, tt.wantErr, err)
			})
		}
	})
}
