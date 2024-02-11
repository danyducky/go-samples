package main

import "fmt"

func main() {
	var a, b int

	fmt.Scan(&a, &b)

	var highest int
	var highestFound bool
	for b >= a {

		if b%7 == 0 {
			highest = b
			highestFound = true
			break
		}

		b--
	}

	if highestFound {
		fmt.Println(highest)
	} else {
		fmt.Println("NO")
	}
}
