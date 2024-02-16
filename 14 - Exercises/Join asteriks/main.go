package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string

	fmt.Scan(&str)

	arr := strings.Split(str, "")
	res := strings.Join(arr, "*")

	fmt.Println(res)
}
