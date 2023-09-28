package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ChangePositions(t *testing.T) {
	h := 44
	b := 77

	ChangePositions(&h, &b)

	assert.Equal(t, 44, b)
	assert.Equal(t, 77, h)

	ChangePositionsGolangStyle(&h, &b)

	assert.Equal(t, 77, b)
	assert.Equal(t, 44, h)
}
