package main

import "fmt"

// Shape adalah interface yang bisa digunakan untuk melakukan perhitungan luas dan keliling untuk sebuah bentuk.
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rect adalah struct untuk menyimpan bentuk lingkaran.
type Rect struct {
	width  float64
	height float64
}

// Area mengimplementasikan interface untuk menghitung luas pada struct Circle.
func (r Rect) Area() float64 {
	return r.width * r.height
}

// DisplayArea mencetak luas sebuah kotak
func DisplayArea(w, h float64) {
	k := Rect{w, h}
	fmt.Printf("luas segiempat adalah %f\n", k.Area())
}
