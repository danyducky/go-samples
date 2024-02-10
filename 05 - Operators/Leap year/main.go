package main

import "fmt"

func main() {
	var year int

	fmt.Scan(&year)

	switch {
	case year%400 == 0:
		fmt.Println("YES")
	case year%4 == 0 && year%100 != 0:
		fmt.Println("YES")
	default:
		fmt.Println("NO")
	}
}
