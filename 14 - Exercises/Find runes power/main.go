package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var number int

	fmt.Scan(&number)

	var arr []string
	var numbers = split(number)
	for _, num := range numbers {
		arr = append(arr, strconv.Itoa(num*num))
	}

	res := strings.Join(arr, "")

	fmt.Println(res)
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
