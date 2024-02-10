package main

import "fmt"

func main() {
	var a float64

	fmt.Scan(&a)

	switch {
	case a <= 0:
		fmt.Printf("число %2.2f не подходит", a)
	case a > 10000:
		fmt.Printf("%e\n", a)
	default:
		fmt.Printf("%.4f", a*a)
	}
}
