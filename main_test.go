package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckVisibility(t *testing.T) {
	t1 := Triangle{
		p1: Point{2.0, 3.6, 4.2},
		p2: Point{1.4, 0.9, 5.3},
		p3: Point{4.8, 8.4, 6.4},
	}

	p := Point{9.6, 16.8, 12.8}

	assert.Equal(t, false, isVisible(p, t1))
}

func TestDet(t *testing.T) {
	m1 := Matrix{
		{1, 0, 4, 6},
		{2, 5, 0, 3},
		{-1, 2, 3, 5},
		{2, 1, -2, 3},
	}

	m2 := Matrix{
		{0, -2, -5, 8},
		{0, -6, 3, 1},
		{0, 11, 5, -3},
		{-1, -2, -1, 3},
	}

	assert.Equal(t, 246.0, Det(m1))
	assert.Equal(t, -441.0, Det(m2))
}
