package geometry

import (
	"math"
)

type Point struct {
	X float64
	Y float64
}

type MultiPoint struct {
	X []float64
	Y []float64
}

type ConnectedMultiPoint interface {
	Coords() (x []float64, y []float64)
}

type Line struct {
	coordinates MultiPoint
}

type Polygon struct {
	coordinates MultiPoint
}

func (p *Line) Coords() ([]float64, []float64) {
	return p.coordinates.X, p.coordinates.Y
}

func (p *Polygon) Coords() ([]float64, []float64) {
	return p.coordinates.X, p.coordinates.Y
}

func _lenatleast2(a []float64) {
	if len(a) <= 1 {
		panic("slice length < 2")
	}
}

func Minf(a []float64) float64 {
	_lenatleast2(a)
	m := a[0]
	for _, v := range a[1:] {
		if v < m {
			m = v
		}
	}
	return m
}

func Maxf(a []float64) float64 {
	_lenatleast2(a)
	m := a[0]
	for _, v := range a[1:] {
		if v > m {
			m = v
		}
	}
	return m
}

func (p *MultiPoint) Bbox() [4]float64 {
	res := [4]float64{Minf(p.X),
		Maxf(p.X),
		Minf(p.Y),
		Maxf(p.Y)}
	return res
}

func (p *MultiPoint) Length() int {
	return len(p.X)
}

func (p *Line) Bbox() [4]float64 {
	return p.coordinates.Bbox()
}

func (p *Polygon) Bbox() [4]float64 {
	return p.coordinates.Bbox()
}

func (p *Line) Length() int {
	return p.coordinates.Length()
}

func (p *Polygon) Length() int {
	return p.coordinates.Length()
}

func (p *Line) Vertex(i int) Point {
	return Point{p.coordinates.X[i], p.coordinates.Y[i]}
}

func (p *Polygon) Vertex(i int) Point {
	return Point{p.coordinates.X[i], p.coordinates.Y[i]}
}

func Cross(x0, x1, y0, y1 float64) float64 {
	return x0*y1 - x1*y0
}

// Return whether a three vertex line turns "counter clockwise"
func CCW(x0, x1, x2, y0, y1, y2 float64) float64 {
	return Cross(x1-x0, x2-x0, y1-y0, y2-y0)
}

// Return true if two segments have a "proper" intersection
func (p *Line) IntersectsProper(q Line) bool {
	px, py := p.Coords()
	qx, qy := q.Coords()
	x0, x1 := px[0], px[1]
	y0, y1 := py[0], py[1]
	x2, x3 := qx[0], qx[1]
	y2, y3 := qy[0], qy[1]
	return CCW(x0, x1, x2, y0, y1, y2)*CCW(x0, x1, x3, y0, y1, y3) < 0 &&
		CCW(x2, x3, x0, y2, y3, y0)*CCW(x2, x3, x1, y2, y3, y1) < 0
}

func (p *Line) Intersects(q Line) bool {
	panic("Not implemented")
}

// Use a crossing number algorithm to compute whether p contains q
func (p *Polygon) Contains(q Point) bool {
	y0 := q.Y
	x0 := Minf(p.coordinates.X) - 1.0
	ray := Line{MultiPoint{[]float64{x0, q.X}, []float64{y0, q.Y}}}

	var cnt int
	var x1, x2, y1, y2 float64
	var seg Line
	px, py := p.Coords()

	n := p.Length()
	for i := 0; i != n; i++ {

		x1 = px[i]
		y1 = py[i]
		if i != n-1 {
			x2 = px[i+1]
			y2 = py[i+1]
		} else {
			x2 = px[0]
			y2 = py[0]
		}
		seg = Line{MultiPoint{[]float64{x1, x2}, []float64{y1, y2}}}

		if ray.IntersectsProper(seg) {
			cnt++
		}
	}
	return math.Mod(float64(cnt), 2) == 1
}

// Test whether any vertex of p (p) is in q (p)
func (p *Polygon) Overlaps(q Polygon) bool {
	return p._overlaps(q) || q._overlaps(*p)
}

// One-sided overlaps test
func (p *Polygon) _overlaps(q Polygon) bool {
	for i := 0; i != q.Length(); i++ {
		if p.Contains(q.Vertex(i)) {
			return true
		}
	}
	return false
}
