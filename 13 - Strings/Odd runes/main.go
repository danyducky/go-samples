package main

import "fmt"

func main() {
	var str string

	fmt.Scan(&str)

	var runes = []rune(str)
	var res string
	for i, r := range runes {
		if i%2 != 0 {
			res += string(r)
		}
	}

	fmt.Println(res)
}
