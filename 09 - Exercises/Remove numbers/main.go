package main

import "fmt"

func main() {
	var number, exclude int

	fmt.Scan(&number, &exclude)

	var arr []int
	for number > 0 {
		digit := number % 10

		if digit != exclude {
			arr = append(arr, digit)
		}

		number = number / 10
	}

	for i := len(arr) - 1; i >= 0; i-- {
		fmt.Print(arr[i])
	}
}
