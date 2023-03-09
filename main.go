package main

import (
	"fmt"
	"log"
)

type Point struct {
	x, y, z float64
}

type Triangle struct {
	p1, p2, p3 Point
}

func main() {
	info()
APP:
	for {
		// given triangle
		var t Triangle
		// given point to check
		var p Point
		var err error

		p, t, err = input()
		if err != nil {
			log.Printf("\nError: %s\n", err.Error())
		} else {
			solve(p, t)
		}

		fmt.Printf("\nDo you want to continue?\nPress [y] - YES; any button - NO\n")
		var s string
		_, _ = fmt.Scan(&s)
		switch s {
		case "y":
			continue
		default:
			break APP
		}
	}
}

func solve(p Point, t Triangle) {
	if isVisible(p, t) {
		fmt.Println("Point is seen")
	} else {
		fmt.Println("Point is not seen")
	}
}

func isVisible(p Point, t Triangle) bool {
	// Unless at least one point of triangle lies
	// in the same octant with given test point, the point can be seen
	if !isSameOctant(p, t) {
		return true
	}

	// Observer's coordinates
	var p0 = Point{x: 0, y: 0, z: 0}

	// If the point is inside the tetrahedron
	// which has vertices p0 and given triangle coords, the point can be seen
	whichTetra := 1
	if isPointInsideOfTetrahedron(p, p0, t, whichTetra) {
		return true
	}

	// The triangle on the same plane with the test point
	tx := ParallelTriangle(p, t)

	// If the point is inside the tetrahedron
	// which has vertices p0 and found triangle coords(which are on a same plane with the test point),
	// the point cannot be seen
	whichTetra = 2
	if isPointInsideOfTetrahedron(p, p0, tx, whichTetra) {
		return false
	}

	// In other cases, the point can be seen
	return true
}

func isPointInsideOfTetrahedron(p, p0 Point, t Triangle, whichTetra int) bool {
	d0 := det(t.p1, t.p2, t.p3, p0)
	d1 := det(p, t.p2, t.p3, p0)
	d2 := det(t.p1, p, t.p3, p0)
	d3 := det(t.p1, t.p2, p, p0)
	d4 := det(t.p1, t.p2, t.p3, p)

	// If the point lies on the given triangle,
	// it cannot be seen
	if whichTetra == 1 && d4 == 0 {
		return false
	}

	if d0*d1 < 0 || d0*d2 < 0 || d0*d3 < 0 || d0*d4 < 0 {
		return false
	}
	return true
}

func isSameOctant(p Point, t Triangle) bool {
	if (t.p1.x*p.x >= 0 && t.p1.y*p.y >= 0 && t.p1.z*p.z >= 0) ||
		(t.p2.x*p.x >= 0 && t.p2.y*p.y >= 0 && t.p2.z*p.z >= 0) ||
		(t.p3.x*p.x >= 0 && t.p3.y*p.y >= 0 && t.p3.z*p.z >= 0) {
		return true
	}
	return false
}

func input() (p Point, t Triangle, err error) {
	fmt.Println("Enter coordinates of triangle:")

	fmt.Print("x1 y1 z1 => ")
	if err = scanf("%f %f %f\n", &t.p1.x, &t.p1.y, &t.p1.z); err != nil {
		return
	}

	fmt.Print("x2 y2 z2 => ")
	if err = scanf("%f %f %f\n", &t.p2.x, &t.p2.y, &t.p2.z); err != nil {
		return
	}

	fmt.Print("x3 y3 z3 => ")
	if err = scanf("%f %f %f\n", &t.p3.x, &t.p3.y, &t.p3.z); err != nil {
		return
	}

	fmt.Println("Enter coordinates of test point")

	fmt.Print("x y z => ")
	if err = scanf("%f %f %f\n", &p.x, &p.y, &p.z); err != nil {
		return
	}

	return
}
