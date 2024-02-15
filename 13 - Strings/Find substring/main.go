package main

import (
	"fmt"
	"strings"
)

func main() {
	var str, substr string

	fmt.Scan(&str, &substr)

	idx := strings.Index(str, substr)

	fmt.Println(idx)
}
