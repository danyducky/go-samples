package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string

	fmt.Scan(&str)

	var runes = []rune(str)
	for _, r := range runes {
		letter := string(r)
		if strings.Count(str, letter) > 1 {
			str = strings.ReplaceAll(str, letter, "")
		}
	}

	fmt.Println(str)
}
