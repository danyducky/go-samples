package main

import "fmt"

func main() {
	var a, b float64

	fmt.Scan(&a, &b)

	avg := (a + b) / 2

	fmt.Println(avg)
}
