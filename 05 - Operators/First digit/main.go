package main

import "fmt"

func main() {
	var a int

	fmt.Scan(&a)

	delim := 1

	switch {
	case a >= 10000:
		delim = 10000
	case a >= 1000:
		delim = 1000
	case a >= 100:
		delim = 100
	case a >= 10:
		delim = 10
	}

	fmt.Println(a / delim)
}
