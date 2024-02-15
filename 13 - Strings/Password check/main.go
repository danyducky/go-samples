package main

import (
	"fmt"
	"unicode"
)

func main() {
	var password string

	fmt.Scan(&password)

	var runes = []rune(password)
	var isWrong bool
	if len(runes) < 5 {
		isWrong = true
	}

	for _, r := range runes {
		switch {
		case unicode.IsPunct(r):
			isWrong = true
			break
		case unicode.IsSymbol(r):
			isWrong = true
			break
		case unicode.IsLetter(r) && !unicode.Is(unicode.Latin, r):
			isWrong = true
			break
		}
	}

	if isWrong {
		fmt.Println("Wrong password")
	} else {
		fmt.Println("Ok")
	}
}
