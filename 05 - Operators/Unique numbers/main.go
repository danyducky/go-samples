package main

import (
	"fmt"
)

func main() {
	var a int

	fmt.Scan(&a)

	numbers := split(a)
	isUnique := getIsUniqueNumbers(numbers)

	switch isUnique {
	case true:
		fmt.Println("YES")
	case false:
		fmt.Println("NO")
	}
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

func getIsUniqueNumbers(numbers []int) bool {
	for i := range numbers {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i] == numbers[j] {
				return false
			}
		}
	}
	return true
}
