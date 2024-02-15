package main

import "fmt"

func main() {
	var str string

	fmt.Scan(&str)

	runes := []rune(str)

	var isPalindrome = true
	for i := 0; i < len(runes); i++ {
		if runes[i] != runes[len(runes)-i-1] {
			isPalindrome = false
			break
		}
	}

	if isPalindrome {
		fmt.Println("Палиндром")
	} else {
		fmt.Println("Нет")
	}
}
