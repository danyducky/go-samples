package main

import "fmt"

func main() {
	var a, b, c int

	fmt.Scan(&a, &b, &c)

	switch {
	case a+b > c && a+c > b && b+c > a:
		fmt.Println("Существует")
	default:
		fmt.Println("Не существует")
	}
}
