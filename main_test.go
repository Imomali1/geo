package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckVisibility(t *testing.T) {
	t1 := Triangle{
		p1: Point{1, 0, 0},
		p2: Point{0, 1, 0},
		p3: Point{0, 0, 1},
	}

	tp1 := Point{1.0 / 3, 1.0 / 3, 1.0 / 3}
	assert.Equal(t, true, isVisible(tp1, t1))

	tp2 := Point{0.5, 0.5, 0.5}
	assert.Equal(t, false, isVisible(tp2, t1))

	tp3 := Point{1, 1, 1}
	assert.Equal(t, false, isVisible(tp3, t1))

	tp4 := Point{-1, -1, -1}
	assert.Equal(t, true, isVisible(tp4, t1))

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
