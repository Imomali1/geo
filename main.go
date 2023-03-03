package main

type Point struct {
	x, y, z float64
}

type Triangle struct {
	p1, p2, p3 Point
}

func main() {
	defer writer.Flush()

	// given triangle
	var t Triangle
	// given point to check
	var p Point

	scanf("%f %f %f\n", &t.p1.x, &t.p1.y, &t.p1.z)
	scanf("%f %f %f\n", &t.p2.x, &t.p2.y, &t.p2.z)
	scanf("%f %f %f\n", &t.p3.x, &t.p3.y, &t.p3.z)
	scanf("%f %f %f\n", &p.x, &p.y, &p.z)

	if isVisible(p, t) {
		printf("%s\n", "Point is seen")
	} else {
		printf("%s\n", "Point is not seen")
	}
}

func isVisible(p Point, t Triangle) (visible bool) {
	tx := ParallelTriangle(p, t)
	// Observer's coordinates
	var p0 = Point{
		x: 0,
		y: 0,
		z: 0,
	}

	d0 := det(tx.p1, tx.p2, tx.p3, p0)
	d1 := det(p, tx.p2, tx.p3, p0)
	d2 := det(tx.p1, p, tx.p3, p0)
	d3 := det(tx.p1, tx.p2, p, p0)
	d4 := det(tx.p1, tx.p2, tx.p3, p)

	if (d1 == 0 || d2 == 0 || d3 == 0 || d4 == 0) ||
		(d0*d1 < 0 || d0*d2 < 0 || d0*d3 < 0 || d0*d4 < 0) {
		visible = false
	} else {
		visible = true
	}

	return
}
