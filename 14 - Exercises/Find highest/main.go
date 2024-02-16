package main

import (
	"fmt"
)

func main() {
	var str string

	fmt.Scan(&str)

	runes := []rune(str)

	var highest rune
	for _, el := range runes {
		if el > highest {
			highest = el
		}
	}

	fmt.Println(string(highest))
}
