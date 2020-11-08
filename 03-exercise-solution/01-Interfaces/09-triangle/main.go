package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}
type triangle struct {
	// lengths of the sides of a triangle.
	a, b, c float64
}
type rectangle struct {
	a, b float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// Heron's Formula for the area of a triangle
func (t triangle) area() float64 {
	p := (t.a + t.b + t.c) / 2.0 // perimeter half
	return math.Sqrt(p * (p - t.a) * (p - t.b) * (p - t.c))
}

func (t rectangle) area() float64 {
	return t.a * t.b
}

func (t triangle) angles() []float64 {
	return []float64{angle(t.b, t.c, t.a), angle(t.a, t.c, t.b), angle(t.a, t.b, t.c)}
}
func angle(a, b, c float64) float64 {
	return math.Acos((a*a+b*b-c*c)/(2*a*b)) * 180.0 / math.Pi
}

type shape interface {
	area() float64
}

func main() {
	shapes := []shape{
		circle{1.0},
		rectangle{5, 10},
		triangle{10, 4, 7},
	}
	for _, v := range shapes {
		fmt.Println(v, "\tArea:", v.area())
		if t, ok := v.(triangle); ok {
			fmt.Println("Angles:", t.angles())
		}
	}
}
