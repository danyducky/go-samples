package main

import "fmt"

func main() {
	var a, b = 1, 5

	swap(&a, &b)
}

func swap(x1 *int, x2 *int) {
	*x1, *x2 = *x2, *x1

	fmt.Println(*x1, *x2)
}
