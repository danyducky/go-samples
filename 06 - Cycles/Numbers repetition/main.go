package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var a, b int

	fmt.Scanln(&a, &b)

	aNums := split(a)
	bNums := split(b)

	var nums []string
	for i := range aNums {
		for j := range bNums {
			if aNums[i] == bNums[j] {
				nums = append(nums, strconv.Itoa(aNums[i]))
			}
		}
	}

	fmt.Println(strings.Join(nums, " "))
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
