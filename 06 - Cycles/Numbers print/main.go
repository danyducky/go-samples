package main

import "fmt"

func main() {
	var number int
	for fmt.Scan(&number); number != 0; fmt.Scan(&number) {
		if number < 10 {
			continue
		}

		if number > 100 {
			break
		}

		fmt.Println(number)
	}
}
