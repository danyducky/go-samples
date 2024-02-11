package main

import "fmt"

func main() {
	var count int

	fmt.Scan(&count)

	var counter int
	for i := 0; i < count; i++ {
		var number int

		fmt.Scan(&number)

		if number > 0 {
			counter++
		}
	}

	fmt.Println(counter)
}
