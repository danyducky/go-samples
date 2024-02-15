package main

import "fmt"

func main() {
	var a, b = 2, 3

	multiply(&a, &b)
}

// считайте что fmt уже импортирован и main объявлен
func multiply(x1 *int, x2 *int) {
	fmt.Println(*x1 * *x2)
}
