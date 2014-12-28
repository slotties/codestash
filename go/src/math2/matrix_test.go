package math2

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMultiply(t *testing.T) {
	m1 := [][]int64 {
		[]int64 { 1, 2 },
		[]int64 { 3, 4 },
	}
	m2 := [][]int64 {
		[]int64 { 5, 6 },
		[]int64 { 7, 8 },
	}

	Matrix2x2Multiply(m1, m2)

	assert.Equal(t, 19, m1[0][0])
	assert.Equal(t, 22, m1[0][1])
	assert.Equal(t, 43, m1[1][0])
	assert.Equal(t, 50, m1[1][1])
}