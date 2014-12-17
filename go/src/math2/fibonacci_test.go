package math2

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestFunctionality(t *testing.T) {
	assert.Equal(t, 0, Fibonacci(-1))
	assert.Equal(t, 0, Fibonacci(0))
	assert.Equal(t, 1, Fibonacci(1))
	assert.Equal(t, 1, Fibonacci(2))
	assert.Equal(t, 2, Fibonacci(3))
	assert.Equal(t, 3, Fibonacci(4))
	assert.Equal(t, 5, Fibonacci(5))
	assert.Equal(t, 8, Fibonacci(6))
	assert.Equal(t, 13, Fibonacci(7))
	assert.Equal(t, 21, Fibonacci(8))
	assert.Equal(t, 34, Fibonacci(9))
	assert.Equal(t, 55, Fibonacci(10))
	assert.Equal(t, 267914296, Fibonacci(42))
	assert.Equal(t, 12586269025, Fibonacci(50))
}

func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibonacci(70000)
	}
}