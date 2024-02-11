package main

import "fmt"

func main() {
	var count int

	fmt.Scan(&count)

	for i := 0; i < count; i++ {
		var number int

		fmt.Scan(&number)

		if i%2 == 0 {
			fmt.Print(number, " ")
		}
	}
}
