package main

import "fmt"

func main() {
	var count int

	fmt.Scan(&count)

	var array = make([]int, count)
	for i := range array {
		var number int

		fmt.Scan(&number)

		array[i] = number
	}

	fmt.Println(array[3])
}
