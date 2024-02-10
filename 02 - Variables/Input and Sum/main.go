package main

import "fmt"

func main() {
	var a, b int

	fmt.Scan(&a) // read 'a'
	fmt.Scan(&b) // read 'b'

	a *= a
	b *= b

	fmt.Println(a + b)
}
