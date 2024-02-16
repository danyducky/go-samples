package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func main() {
	var num = adding("%^80", "hhhhh20&&&&nd")

	fmt.Println(num)
}

func adding(a string, b string) int64 {
	a = clear(a)
	b = clear(b)

	aNum, _ := strconv.ParseInt(a, 10, 0)
	bNum, _ := strconv.ParseInt(b, 10, 0)

	return aNum + bNum
}

func clear(str string) string {
	var runes = []rune(str)
	var res string
	for _, r := range runes {
		if unicode.IsDigit(r) {
			res += string(r)
		}
	}
	return res
}
