package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetIbit(t *testing.T) {

	t.Run("power of two", func(t *testing.T) {

		expecterPowerOfTwo := int64(1)
		for i := 0; i < 64; i++ {
			number := new(int64)
			SetIbit(number, i, 1)

			assert.Equal(t, expecterPowerOfTwo, *number)
			expecterPowerOfTwo *= 2
		}
	})

	t.Run("change sign", func(t *testing.T) {
		max := int64(0x7FFF_FFFF_FFFF_FFFF)
		FFFF_FFFF_FFFF_FFFF := int64(-1)
		SetIbit(&max, 63, 1)
		assert.Equal(t, FFFF_FFFF_FFFF_FFFF, max)

		SetIbit(&max, 63, 0)
		assert.Equal(t, int64(0x7FFF_FFFF_FFFF_FFFF), max)
	})

	t.Run("set zeros", func(t *testing.T) {
		ref := int64(0b111_1111_1111_1111_1111_1111_1111_1111_1111_1111_1111_1111_1111_1111_1111_1111)
		z60 := int64(0b110_1111_1111_1111_1111_1111_1111_1111_1111_1111_1111_1111_1111_1111_1111_1111)
		z45and60 := int64(0b110_1111_1111_1111_1101_1111_1111_1111_1111_1111_1111_1111_1111_1111_1111_1111)
		z30and45and60 := int64(0b110_1111_1111_1111_1101_1111_1111_1111_1011_1111_1111_1111_1111_1111_1111_1111)
		z15and30and45and60 := int64(0b110_1111_1111_1111_1101_1111_1111_1111_1011_1111_1111_1111_0111_1111_1111_1111)
		z0and15and30and45and60 := int64(0b110_1111_1111_1111_1101_1111_1111_1111_1011_1111_1111_1111_0111_1111_1111_1110)

		SetIbit(&ref, 60, 0)
		assert.Equal(t, z60, ref)

		SetIbit(&ref, 45, 0)
		assert.Equal(t, z45and60, ref)

		SetIbit(&ref, 30, 0)
		assert.Equal(t, z30and45and60, ref)

		SetIbit(&ref, 15, 0)
		assert.Equal(t, z15and30and45and60, ref)

		SetIbit(&ref, 0, 0)
		assert.Equal(t, z0and15and30and45and60, ref)

	})
}
