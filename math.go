package main

type Matrix [][]float64
type stack []float64

func ParallelTriangle(p Point, t Triangle) Triangle {
	var newT Triangle

	x1, y1, z1 := t.p1.x, t.p1.y, t.p1.z
	x2, y2, z2 := t.p2.x, t.p2.y, t.p2.z
	x3, y3, z3 := t.p3.x, t.p3.y, t.p3.z

	a := (y2-y1)*(z3-z1) - (y3-y1)*(z2-z1)
	b := (x3-z1)*(z2-z1) - (x2-x1)*(z3-z1)
	c := (x2-x1)*(y3-y1) - (x3-x1)*(y2-y1)
	d := -a*p.x - b*p.y - c*p.z

	k := (-d) / (a*x1 + b*y1 + c*z1)
	newT.p1.x, newT.p1.y, newT.p1.z = k*x1, k*y1, k*z1

	k = (-d) / (a*x2 + b*y2 + c*z2)
	newT.p2.x, newT.p2.y, newT.p2.z = k*x2, k*y2, k*z2

	k = (-d) / (a*x3 + b*y3 + c*z3)
	newT.p3.x, newT.p3.y, newT.p3.z = k*x3, k*y3, k*z3

	return newT
}

func det(p1, p2, p3, p4 Point) float64 {
	matrix := Matrix{
		{p1.x, p1.y, p1.z, 1},
		{p2.x, p2.y, p2.z, 1},
		{p3.x, p3.y, p3.z, 1},
		{p4.x, p4.y, p4.z, 1},
	}
	d := Det(matrix)
	return d
}

func Det(mat Matrix) float64 {
	if len(mat) == 2 {
		return (mat[0][0] * mat[1][1]) - (mat[0][1] * mat[1][0])
	}

	s := 0.0

	for i := 0; i < len(mat[0]); i++ {
		sm := subMat(mat[1:][:], i)
		z := Det(sm)

		if i%2 != 0 {
			s -= mat[0][i] * z
		} else {
			s += mat[0][i] * z
		}
	}

	return s
}

func subMat(mat Matrix, p int) Matrix {
	stacks := make([]stack, len(mat))
	for n := range mat {
		stacks[n] = stack{}
		for j := range mat[n] {
			if j != p {
				stacks[n].push(mat[n][j])
			}
		}
	}
	out := make(Matrix, len(mat))
	for k := range stacks {
		out[k] = stacks[k].ToSlice()
	}
	return out
}

func (s *stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *stack) push(n float64) {
	*s = append(*s, n)
}

func (s *stack) pop() (float64, bool) {
	if s.isEmpty() {
		return 0, false
	}
	i := len(*s) - 1
	n := (*s)[i]
	*s = (*s)[:i]
	return n, true
}

func (s *stack) ToSlice() []float64 {
	return *s
}
