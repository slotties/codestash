package math2

func Fibonacci(n int) int64 {
	// return fibonacciMatrix(n)
	return fibonacciIterative(n)
}

// This is horribly slow.
func fibonacciRecursive(n int) int {
	if n < 2 {
    	return n
	}
	return fibonacciRecursive(n - 1) + fibonacciRecursive(n - 2)
}
// This is kind of fast actually.
func fibonacciIterative(n int) int64 {
	if n < 1 {
		return 0
	}

	prev := int64(0)
	current := int64(1)
	for i := 1; i < n; i++ {
		next := prev + current
		prev = current
		current = next
	}

	return current
}

// http://kukuruku.co/hub/algorithms/the-nth-fibonacci-number-in-olog-n
func fibonacciMatrix(n int) int64 {
	if n <= 0 {
		return 0
	}

	matrix := [][]int64 {
		[]int64 { 1, 1 },
		[]int64 { 1, 0 },
	}

	Matrix2x2Power(matrix, n - 1)

	return matrix[0][0]
}