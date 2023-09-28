package main

import (
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_mySleep(t *testing.T) {
	tests := []time.Duration{time.Second, time.Microsecond, time.Millisecond, 500 * time.Millisecond, 150 * time.Millisecond}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			startTime := time.Now()
			mySleep(tt)
			d := time.Since(startTime)

			assert.GreaterOrEqual(t, d, tt)
		})
	}
}
