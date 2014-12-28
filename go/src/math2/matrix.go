package math2

func Matrix2x2Power(matrix [][]int64, power int) {
	powerMatrix := [][]int64 {
		[]int64 { 1, 1 },
		[]int64 { 1, 0 },
	}

	for i := 2; i <= power; i++ {
		Matrix2x2Multiply(matrix, powerMatrix)
	}
}

func Matrix2x2Multiply(m1 [][]int64, m2 [][]int64) {
	a := m1[0][0] * m2[0][0] + m1[0][1] * m2[1][0]
	b := m1[0][0] * m2[0][1] + m1[0][1] * m2[1][1]
	c := m1[1][0] * m2[0][0] + m1[1][1] * m2[1][0]
	d := m1[1][0] * m2[0][1] + m1[1][1] * m2[1][1]

	m1[0][0] = a
	m1[0][1] = b
	m1[1][0] = c
	m1[1][1] = d
}