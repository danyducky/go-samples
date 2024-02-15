package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	text = strings.TrimSuffix(text, "\n")

	runes := []rune(text)
	letter := runes[0]
	dot := runes[len(runes)-1]

	isLetter := unicode.IsLetter(letter)
	isCapitalized := unicode.IsUpper(letter)
	if isLetter && isCapitalized && dot == '.' {
		fmt.Println("Right")
	} else {
		fmt.Println("Wrong")
	}
}
