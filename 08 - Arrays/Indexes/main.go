package main

import "fmt"

func main() {
	var workArray [10]uint8

	for i := 0; i < 10; i++ {
		var number uint8

		fmt.Scan(&number)

		workArray[i] = number
	}

	for j := 0; j < 6; j += 2 {
		var a, b uint8

		fmt.Scan(&a, &b)

		workArray[a], workArray[b] = workArray[b], workArray[a]
	}

	for _, el := range workArray {
		fmt.Printf("%d ", el)
	}
}
