package main

import "fmt"

func main() {
	var a, b, c int

	fmt.Scan(&a, &b, &c)

	a *= a
	b *= b
	c *= c

	switch {
	case a+b == c:
		fmt.Println("Прямоугольный")
	default:
		fmt.Println("Непрямоугольный")
	}
}
