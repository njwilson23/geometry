package geometry

import (
	"testing"
)

func TestCreatePoint(t *testing.T) {
	pt := Point{2, 3}
	if (pt.X) != 2 || (pt.Y != 3) {
		t.Fail()
	}
}

func TestCreateMultiPoint(t *testing.T) {
	mp := MultiPoint{[]float64{3, 4, 5, 6}, []float64{8, 6, 4, 2}}
	if (mp.X[0] != 3) || (mp.Y[1] != 6) || (mp.X[3] != 6) {
		t.Fail()
	}
}

func TestCreateLine(t *testing.T) {
	mp := MultiPoint{[]float64{3, 4, 5, 6}, []float64{8, 6, 4, 2}}
	line := Line{mp}
	if (line.coordinates.X[0] != 3) ||
		(line.coordinates.Y[1] != 6) ||
		(line.coordinates.X[3] != 6) {
		t.Fail()
	}
}

func TestBbox(t *testing.T) {
	mp := MultiPoint{[]float64{3, 4, 5, 6}, []float64{8, 6, 4, 2}}
	if mp.Bbox() != [4]float64{3, 6, 2, 8} {
		t.Fail()
	}
}

func TestLength(t *testing.T) {
	p := Line{MultiPoint{[]float64{3, 4, 5, 6}, []float64{8, 6, 4, 2}}}
	if p.Length() != 4 {
		t.Fail()
	}
}

func TestVertex(t *testing.T) {
	p := Polygon{MultiPoint{[]float64{3, 4, 5, 6}, []float64{8, 6, 4, 2}}}
	q := Point{5, 4}
	if p.Vertex(2) != q {
		t.Fail()
	}
}

func TestIntersects(t *testing.T) {
	line1 := Line{MultiPoint{[]float64{0, 1}, []float64{0, 1}}}
	line2 := Line{MultiPoint{[]float64{0, 1}, []float64{1, 0}}}
	if !line1.Intersects(line2) {
		t.Fail()
	}

	// Crossing the left vertex counts but the right doesn't
	line3 := Line{MultiPoint{[]float64{-1, 1}, []float64{0, 0}}}
	line4 := Line{MultiPoint{[]float64{0, 2}, []float64{1, 1}}}
	if !line1.Intersects(line3) {
		t.Fail()
	}
	if line1.Intersects(line4) {
		t.Fail()
	}
}

func TestIntersectsProper(t *testing.T) {
	line1 := Line{MultiPoint{[]float64{0, 1}, []float64{0, 1}}}
	line2 := Line{MultiPoint{[]float64{0, 1}, []float64{1, 0}}}
	if !line1.IntersectsProper(line2) {
		t.Fail()
	}

	// Crossing vertices doesn't count
	line3 := Line{MultiPoint{[]float64{-1, 1}, []float64{0, 0}}}
	line4 := Line{MultiPoint{[]float64{0, 2}, []float64{1, 1}}}
	if line1.IntersectsProper(line3) {
		t.Fail()
	}
	if line1.IntersectsProper(line4) {
		t.Fail()
	}
}

func TestContains(t *testing.T) {
	p := Polygon{MultiPoint{[]float64{0, 1, 1, 0}, []float64{0, 0, 1, 1}}}
	pt := Point{0.5, 0.5}
	if !p.Contains(pt) {
		t.Fail()
	}

	pt2 := Point{2, 1}
	if p.Contains(pt2) {
		t.Fail()
	}
}
