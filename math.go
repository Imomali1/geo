package main

type Matrix [][]float64
type stack []float64

func ParallelTriangle(p Point, t Triangle) Triangle {
	var t1 Triangle

	a := (t.p2.y-t.p1.y)*(t.p3.z-t.p1.z) - (t.p3.y-t.p1.y)*(t.p2.z-t.p1.z)
	b := (t.p3.x-t.p1.z)*(t.p2.z-t.p1.z) - (t.p2.x-t.p1.x)*(t.p3.z-t.p1.z)
	c := (t.p2.x-t.p1.x)*(t.p3.y-t.p1.y) - (t.p3.x-t.p1.x)*(t.p2.y-t.p1.y)
	d := -a*p.x - b*p.y - c*p.z

	k1 := (-d) / (a*t.p1.x + b*t.p1.y + c*t.p1.z)
	t1.p1.x, t1.p1.y, t1.p1.z = k1*t.p1.x, k1*t.p1.y, k1*t.p1.z
	k2 := (-d) / (a*t.p2.x + b*t.p2.y + c*t.p2.z)
	t1.p2.x, t1.p2.y, t1.p2.z = k2*t.p2.x, k2*t.p2.y, k2*t.p2.z
	k3 := (-d) / (a*t.p3.x + b*t.p3.y + c*t.p3.z)
	t1.p3.x, t1.p3.y, t1.p3.z = k3*t.p3.x, k3*t.p3.y, k3*t.p3.z

	return t1
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
