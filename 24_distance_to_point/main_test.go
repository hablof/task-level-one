package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Distance(t *testing.T) {
	tests := []struct {
		x1   float64
		y1   float64
		x2   float64
		y2   float64
		dist float64
	}{
		{
			x1:   0,
			y1:   0,
			x2:   0,
			y2:   0,
			dist: 0,
		},
		{
			x1:   10,
			y1:   0,
			x2:   0,
			y2:   0,
			dist: 10,
		},
		{
			x1:   10,
			y1:   4,
			x2:   4,
			y2:   12,
			dist: 10,
		},
		{
			x1:   0,
			y1:   0.6,
			x2:   0.8,
			y2:   0,
			dist: 1,
		},
	}
	for _, tt := range tests {
		p1 := NewPoint(tt.x1, tt.y1)
		p2 := NewPoint(tt.x2, tt.y2)

		assert.Equal(t, tt.dist, p1.DistanceToPoint(p2))
		assert.Equal(t, tt.dist, p2.DistanceToPoint(p1))
	}
}
