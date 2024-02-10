package main

import "fmt"

func main() {
	var a int

	fmt.Scan(&a)

	numbers := split(a)

	firstThree := sum(numbers[:3])
	lastThree := sum(numbers[3:])

	switch {
	case firstThree == lastThree:
		fmt.Println("YES")
	default:
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

func sum(numbers []int) int {
	var sum int
	for _, number := range numbers {
		sum += number
	}
	return sum
}
