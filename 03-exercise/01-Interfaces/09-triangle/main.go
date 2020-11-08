package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}
type triangle struct {
	a, b, c float64 // lengths of the sides of a triangle.
}
type rectangle struct {
	h, w float64
}

func (c circle) String() string {
	return fmt.Sprint("Circle (Radius: ", c.radius, ")")
}
func (t triangle) String() string {
	return fmt.Sprint("Triangle (Sides: ", t.a, ", ", t.b, ", ", t.c, ")")
}
func (r rectangle) String() string {
	return fmt.Sprint("Rectangle (Sides: ", r.h, ", ", r.w, ")")
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (t triangle) area() float64 {
	// Heron's formula
	p := (t.a + t.b + t.c) / 2.0
	return math.Sqrt(p * (p - t.a) * (p - t.b) * (p - t.c))
}

func (r rectangle) area() float64 {
	return r.h * r.w
}

func (t triangle) angles() []float64 {
	return []float64{angle(t.b, t.c, t.a), angle(t.a, t.c, t.b), angle(t.a, t.b, t.c)}
}

func angle(a, b, c float64) float64 {
	return math.Acos((a*a+b*b-c*c)/(2*a*b)) * 180.0 / math.Pi
}

// TODO: Instantiate circle, triangle, rectangle with values
// Print area of each.
// print angles of the triangle

func main() {

}
