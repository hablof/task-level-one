package main

import (
	"strconv"
	"testing"
)

func Test_determineGroup(t *testing.T) {
	tests := []struct {
		val  float32
		want int
	}{
		{
			val:  0,
			want: 0,
		},
		{
			val:  1.2,
			want: 0,
		},
		{
			val:  11.2,
			want: 10,
		},
		{
			val:  -1.2,
			want: -10,
		},
		{
			val:  -0.0,
			want: 0,
		},
		{
			val:  -0.1,
			want: -10,
		},
		{
			val:  21.2,
			want: 20,
		},
		{
			val:  -21.2,
			want: -30,
		},
	}
	for _, tt := range tests {
		t.Run(strconv.FormatFloat(float64(tt.val), 'f', -1, 32), func(t *testing.T) {
			if got := determineGroup(tt.val); got != tt.want {
				t.Errorf("determineGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}
