package main

import "fmt"

func main() {
	var count int

	fmt.Scanln(&count)

	var sum int
	for i := 0; i < count; i++ {
		var number int

		fmt.Scan(&number)

		if number%8 != 0 {
			continue
		}

		if 9 < number && number <= 99 {
			sum += number
		}
	}

	fmt.Println(sum)
}
