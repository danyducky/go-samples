package main

import "fmt"

func main() {
	var number int

	fmt.Scan(&number)

	numbers := split(number)

	var sum int
	for _, num := range numbers {
		sum += num
	}

	root := split(sum)

	sum = 0
	for _, num := range root {
		sum += num
	}

	fmt.Println(sum)
}

func split(number int) []int {
	var slice []int
	for number > 0 {

		slice = append(slice, number%10)

		number /= 10
	}

	var reverse []int
	for i := range slice {
		reverse = append(reverse, slice[len(slice)-1-i])
	}

	return reverse
}
