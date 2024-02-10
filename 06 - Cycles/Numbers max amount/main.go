package main

import "fmt"

func main() {
	var number int
	var numbersMap = make(map[int]int)

	for fmt.Scan(&number); number != 0; fmt.Scan(&number) {
		numbersMap[number] += 1
	}

	var maxKey int
	for key := range numbersMap {
		if key > maxKey {
			maxKey = key
		}
	}

	fmt.Println(numbersMap[maxKey])
}
