package main

import "fmt"

func main() {
	var count int

	fmt.Scan(&count)

	var array = make([]int, count)
	for i := 0; i < count; i++ {
		var number int

		fmt.Scanln(&number)

		array[i] = number
	}

	var minimum = array[0]
	var counter = 0
	for _, number := range array {
		switch {
		case number == minimum:
			counter++
		case number < minimum:
			minimum = number
			counter = 1
		}
	}

	fmt.Println(counter)
}
