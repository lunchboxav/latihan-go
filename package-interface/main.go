package main

import (
	"fmt"
	"math"
)

// Circle adalah struct untuk menyimpan bentuk lingkaran.
type Circle struct {
	radius float64
}

// Area Mengimplementasi interface untuk menghitung luas pada struct Circle.
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {
	DisplayArea(5, 4)
	l := Circle{10}
	fmt.Printf("luas lingkaran adalah %f\n", l.Area())
}
