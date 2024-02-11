package main

import "fmt"

func main() {
	var number int

	fmt.Scan(&number)

	numbers := splitReverse(number)

	for i := range numbers {
		fmt.Print(numbers[i])
	}

}

func splitReverse(number int) []int {
	var slice []int
	for number > 0 {

		slice = append(slice, number%10)

		number /= 10
	}
	return slice
}
