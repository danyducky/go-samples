package main

import (
	"fmt"
)

func main() {
	var number int

	fmt.Scan(&number)

	var value = 1
	for value <= number {
		fmt.Printf("%d ", value)

		value *= 2
	}
}
