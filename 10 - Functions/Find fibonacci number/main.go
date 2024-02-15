package main

import "fmt"

func main() {
	var number = getFibonacciNumber(4)

	fmt.Println(number)
}

func getFibonacciNumber(n int) int {
	arr := []int{0, 1}
	for i := 0; i < n; i++ {
		arr = append(arr, arr[i]+arr[i+1])
	}
	return arr[len(arr)-1]
}
