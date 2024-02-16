package main

import (
	"fmt"
	"math"
)

type Triangle struct {
	CatA float64
	CatB float64
}

func NewTriangle(catA float64, catB float64) *Triangle {
	return &Triangle{
		CatA: catA,
		CatB: catB,
	}
}

func (triangle *Triangle) GetHypotenuse() float64 {
	a := math.Pow(triangle.CatA, 2)
	b := math.Pow(triangle.CatB, 2)

	return math.Sqrt(a + b)
}

func main() {
	var a, b float64

	fmt.Scan(&a, &b)

	var triangle = NewTriangle(a, b)

	fmt.Println(triangle.GetHypotenuse())
}
