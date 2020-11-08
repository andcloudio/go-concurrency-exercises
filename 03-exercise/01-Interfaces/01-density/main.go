package main

import "fmt"

// Metal - mass and volume information
type Metal struct {
	mass   float64
	volume float64
}

// Density - return density of metal
func (m *Metal) Density() float64 {
	// density = mass/volume
	return m.mass / m.volume
}

// IsDenser - compare density of two objects
func IsDenser(a, b *Metal) bool {
	return a.Density() > b.Density()
}

func main() {
	gold := Metal{478, 24}
	silver := Metal{100, 10}

	result := IsDenser(&gold, &silver)
	if result {
		fmt.Println("gold has higher density than silver")
	} else {
		fmt.Println("silver has higher density than gold")
	}
}
