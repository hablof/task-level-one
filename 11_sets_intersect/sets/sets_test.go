package sets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testCases := []struct {
		desc              string
		elementsSet1      []int
		elementsSet2      []int
		elementsIntersect []int
	}{
		{
			desc:              "",
			elementsSet1:      nil,
			elementsSet2:      []int{},
			elementsIntersect: []int{},
		},
		{
			desc:              "",
			elementsSet1:      []int{1, 2, 3, 4, 5},
			elementsSet2:      []int{6, 7, 8, 9, 10},
			elementsIntersect: []int{},
		},
		{
			desc:              "",
			elementsSet1:      []int{1, 2, 3, 4, 5},
			elementsSet2:      []int{1, 2, 3, 4, 5},
			elementsIntersect: []int{1, 2, 3, 4, 5},
		},
		{
			desc:              "",
			elementsSet1:      []int{0, 1, 2, 3, 5, 7, 11, 13},
			elementsSet2:      []int{0, 2, 4, 6, 8, 10},
			elementsIntersect: []int{0, 2},
		},
		{
			desc:              "",
			elementsSet1:      []int{45, 115, 180},
			elementsSet2:      []int{115, 901, 14},
			elementsIntersect: []int{115},
		},
	}
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			set1 := NewSet(tt.elementsSet1...)
			set2 := NewSet(tt.elementsSet2...)
			set1IntersectsSet2 := set1.Intersect(set2)
			set2IntersectsSet1 := set2.Intersect(set1)

			assert.ElementsMatch(t, tt.elementsIntersect, set1IntersectsSet2.Elements())
			assert.ElementsMatch(t, tt.elementsIntersect, set2IntersectsSet1.Elements())
		})
	}
}
