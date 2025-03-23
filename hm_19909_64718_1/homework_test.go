package main

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestFactorial(t *testing.T) {
	testCases := map[int]struct {
		result int
		err    error
	}{
		3: {
			result: 6,
			err:    nil,
		},
		5: {
			result: 120,
			err:    nil,
		},
		1: {
			result: 1,
			err:    nil,
		},
		0: {
			result: 1,
			err:    nil,
		},
		-1: {
			result: 0,
			err:    ErrNegativeFactorial,
		},
	}

	for x, tc := range testCases {
		t.Run(strconv.Itoa(x), func(t *testing.T) {
			result, err := factorial(x)

			assert.Equal(t, tc.result, result, "testing for (equal results)")
			assert.ErrorIs(t, err, tc.err, "testing for (equal errors)")
		})
	}
}
