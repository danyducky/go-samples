package main

import "fmt"

func main() {
	var a int
	var b int

	fmt.Scan(&a, &b)

	var sum int
	for a <= b {
		sum += a
		a++
	}

	fmt.Println(sum)
}
