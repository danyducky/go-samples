package main

import "fmt"

func main() {
	var x, p, y int

	fmt.Scan(&x, &p, &y)

	var years int
	for x < y {
		x += x * p / 100
		years++
	}

	fmt.Println(years)
}
