package main

import "fmt"

func main() {
	var number int

	fmt.Scan(&number)

	arr := []int{0, 1}
	for i := 0; i < 999; i++ {
		arr = append(arr, arr[i]+arr[i+1])
	}

	idx := -1
	for i, el := range arr {
		if el == number {
			idx = i
		}
	}

	fmt.Println(idx)
}
